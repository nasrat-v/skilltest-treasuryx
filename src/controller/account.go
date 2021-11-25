package controller

import (
	"skilltest-treasuryx/src/database"
)

// Find creditor and debtor account
// If not found, create them
func (x *Controller) findOrCreateCreditorDebtorAccounts(paymentReq PaymentRequest) (database.Account, database.Account, error) {
	creditorAccount, err := x.findOrCreateAccount(database.Account{
		Iban: paymentReq.CreditorIban,
		Name: paymentReq.CreditorName,
	})
	if err != nil {
		return database.Account{}, database.Account{}, err
	}
	debtorAccount, err := x.findOrCreateAccount(database.Account{
		Iban: paymentReq.DebtorIban,
		Name: paymentReq.DebtorName,
	})
	if err != nil {
		return database.Account{}, database.Account{}, err
	}
	return creditorAccount, debtorAccount, nil
}

func (x *Controller) findOrCreateAccount(account database.Account) (database.Account, error) {
	dbAccount, err := x._database.GetAccountByIban(account.Iban)
	if err != nil {
		return account, err
	}
	if dbAccount.Id == 0 {
		return x._database.InsertAccount(database.Account{
			Name: account.Name,
			Iban: account.Iban,
		})
	}
	return dbAccount, nil
}
