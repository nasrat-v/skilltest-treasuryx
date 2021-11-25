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
	databasePath := os.Getenv("SQLITE_DB_FILE_LOCATION")
	err := x.createFile(databasePath)
	if err != nil {
		return err
	}
	x._sqliteDatabase, err = sql.Open("sqlite3", databasePath) // Open the created SQLite File
	if err != nil {
		return err
	}
	return x.createTables()
}

func (x *Database) Close() error {
	return x._sqliteDatabase.Close()
}

// Create database file
// Remove if already existing
func (x *Database) createFile(databasePath string) error {
	if _, err := os.Stat(databasePath); err == nil { // Remove if database exists
		if err := os.Remove(databasePath); err != nil {
			return err
		}
	}
	file, err := os.Create(databasePath) // Create SQLite file
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

// Initialise db tables
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
	statementAccount, err := x._sqliteDatabase.Prepare(createAccountTableSQL)
	if err != nil {
		return err
	}
	statementPayment, err := x._sqliteDatabase.Prepare(createPaymentTableSQL)
	if err != nil {
		return err
	}
	statementAccount.Exec()
	statementPayment.Exec()
	return nil
}
