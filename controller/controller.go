package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"skilltest-treasuryx/database"

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
	}
	// Check if creditor and debtor accounts exist in db. If not, create missing
	creditorAccount, debtorAccount, err := x.findOrCreateCreditorDebtorAccounts(paymentReq)
	if err != nil {
		context.Abort()
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	// Create payment in db with accounts ids
	payment, err := x.createPaymentFromAccounts(creditorAccount.Id, debtorAccount.Id, paymentReq)
	if err != nil {
		context.Abort()
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	fmt.Println(payment)
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
