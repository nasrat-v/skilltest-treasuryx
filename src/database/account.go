package database

import (
	_ "github.com/mattn/go-sqlite3"
)

// Fetch account data from IBAN
func (x *Database) GetAccountByIban(iban string) (Account, error) {
	account := Account{}

	rows, err := x._sqliteDatabase.Query("SELECT id, name, iban FROM account WHERE iban = '" + iban + "'")
	if err != nil {
		return account, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&account.Id, &account.Name, &account.Iban); err != nil {
			return account, err
		}
	}
	return account, nil
}

// Create new account in db
func (x *Database) InsertAccount(account Account) (Account, error) {
	stmt, err := x._sqliteDatabase.Prepare("INSERT INTO account (id, iban, name) VALUES (?, ?, ?)")
	if err != nil {
		return Account{}, err
	}
	if _, err := stmt.Exec(nil, account.Iban, account.Name); err != nil {
		return Account{}, err
	}
	defer stmt.Close()
	return x.GetAccountByIban(account.Iban) // fetch again to get freshly inserted account
}
