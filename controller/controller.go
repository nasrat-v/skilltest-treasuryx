package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"skilltest-treasuryx/bank"
	"skilltest-treasuryx/database"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	_validate *validator.Validate
	_database *database.Database
}

// default constructor
func New() Controller {
	return Controller{}
}

func (x *Controller) Create(database *database.Database) {
	x._validate = validator.New()
	x._database = database
}

func (x *Controller) Payment(context *gin.Context) {
	errorMsg := "Bad Request"

	paymentReq, err := x.decodePaymentRequest(context.Request.Body)
	// Validate body
	errValidator := x._validate.Struct(paymentReq)
	if err != nil || errValidator != nil {
		if errValidator != nil {
			errorMsg = errValidator.Error()
		}
		context.Abort()
		context.JSON(http.StatusBadRequest, gin.H{
			"message": errorMsg,
		})
		return
	}
	// Check if creditor and debtor accounts exist in db. If not, create missing
	creditorAccount, debtorAccount, err := x.findOrCreateCreditorDebtorAccounts(paymentReq)
	if err != nil {
		context.Abort()
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Create payment in db with accounts ids
	payment, err := x.createPaymentFromAccounts(creditorAccount.Id, debtorAccount.Id, paymentReq)
	if err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(err.Error(), "UNIQUE constraint failed") { // check for doublon
			status = http.StatusBadRequest
		}
		context.Abort()
		context.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Create Bank XML File
	err = bank.CreateXmlFile(bank.MarshalDocument(creditorAccount, debtorAccount, payment))
	if err != nil {
		context.Abort()
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": "Payment successfully transmited to the bank",
		})
	}
	// Get Bank response
	status, err := bank.GetBankStatusResponse(payment.IdempotencyUniqueKey)
	if err != nil {
		return
	}
	// Update Payment status in db
	if status != "" {
		_, err := x._database.UpdatePaymentStatusByIdempotency(status, payment.IdempotencyUniqueKey)
		if err != nil {
			return
		}
	}
}

func (x *Controller) decodePaymentRequest(body io.ReadCloser) (PaymentRequest, error) {
	var paymentReq PaymentRequest

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&paymentReq)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return paymentReq, err
	}
	return paymentReq, nil
}
