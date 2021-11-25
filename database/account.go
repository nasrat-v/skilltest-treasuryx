package database

import (
	_ "github.com/mattn/go-sqlite3"
)

func (x *Database) InsertAccount(account Account) error {
	stmt, err := x._sqliteDatabase.Prepare("INSERT INTO account (id, iban, name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	stmt.Exec(nil, account.Iban, account.Name)
	defer stmt.Close()
	return nil
}
