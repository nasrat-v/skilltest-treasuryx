package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	_sqliteDatabase *sql.DB
}

// default constructor
func New() Database {
	return Database{}
}

func (x *Database) Create() error {
	err := x.createFile()
	if err != nil {
		return err
	}
	x._sqliteDatabase, err = sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		return err
	}
	return x.createTables()
}

func (x *Database) Close() error {
	return x._sqliteDatabase.Close()
}

func (x *Database) createFile() error {
	if _, err := os.Stat("sqlite-database.db"); err == nil {
		if err := os.Remove("sqlite-database.db"); err != nil {
			return err
		}
	}
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

func (x *Database) createTables() error {
	createAccountTableSQL := `CREATE TABLE account (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"iban" TEXT UNIQUE NOT NULL,
		"name" TEXT NOT NULL
		);`
	createPaymentTableSQL := `CREATE TABLE payment (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"debtorId" INTEGER NOT NULL,
		"creditorId" INTEGER NOT NULL,
		"ammount" REAL NOT NULL,
		"idempotencyUniqueKey" TEXT UNIQUE NOT NULL,
		"status" TEXT NOT NULL
	);`
	statementAccount, err := x._sqliteDatabase.Prepare(createAccountTableSQL) // Prepare SQL Statement
	if err != nil {
		return err
	}
	statementPayment, err := x._sqliteDatabase.Prepare(createPaymentTableSQL) // Prepare SQL Statement
	if err != nil {
		return err
	}
	statementAccount.Exec()
	statementPayment.Exec()
	return nil
}
