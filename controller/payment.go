package controller

import "skilltest-treasuryx/database"

func (x *Controller) createPaymentFromAccounts(creditorId int, debtorId int, paymentReq PaymentRequest) (database.Payment, error) {
	return x._database.InsertPayment(database.Payment{
		CreditorId:           creditorId,
		DebtorId:             debtorId,
		Ammount:              paymentReq.Ammount,
		IdempotencyUniqueKey: paymentReq.IdempotencyUniqueKey,
		Status:               "CREATED",
	})
}
