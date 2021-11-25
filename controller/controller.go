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

	payment, err := x.decodePaymentRequest(context.Request.Body)
	errValidator := x._validate.Struct(payment)
	if err != nil || errValidator != nil {
		if errValidator != nil {
			errorMsg = errValidator.Error()
		}
		context.Abort()
		context.JSON(http.StatusBadRequest, gin.H{
			"message": errorMsg,
		})
	}
	// check if creditor and debtor account exist and get id
	// if not create it
	x.findOrCreateAccounts(payment)
}

func (x *Controller) findOrCreateAccounts(payment Payment) {
	account := database.Account{
		Name: payment.CreditorName,
		Iban: payment.CreditorIban,
	}
	if err := x._database.InsertAccount(account); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func (x *Controller) decodePaymentRequest(body io.ReadCloser) (Payment, error) {
	var payment Payment

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&payment)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return payment, err
	}
	return payment, nil
}
