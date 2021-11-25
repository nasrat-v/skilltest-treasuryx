package controller

import (
	"skilltest-treasuryx/src/bank"
	"time"
)

func (x *Controller) bankResponseForPayment(idempotency string) error {
	// Get Bank response
	status, fileExist, err := bank.GetBankStatusResponse(idempotency)
	if err != nil {
		return err
	}
	for !fileExist { // Iterate until we found the file
		time.Sleep(time.Second)
		return x.bankResponseForPayment(idempotency) // Recursivity
	}
	// If file exist and status is not empty, update Payment status in db
	if status != "" {
		_, err := x._database.UpdatePaymentStatusByIdempotency(status, idempotency)
		if err != nil {
			return err
		}
	}
	return nil
}
