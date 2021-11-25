package main

import (
	"errors"
	"fmt"
	"os"
	"skilltest-treasuryx/src/manager"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func checkEnv() error {
	bankFolder := os.Getenv("BANK_FOLDER")
	sqliteDbFileLocation := os.Getenv("SQLITE_DB_FILE_LOCATION")

	if bankFolder == "" {
		return errors.New("error: No BANK_FOLDER env variable provided")
	}
	if sqliteDbFileLocation == "" {
		return errors.New("error: No SQLITE_DB_FILE_LOCATION env variable provided")
	}
	return nil
}

func main() {
	godotenv.Load()
	gin.SetMode(os.Getenv("GIN_MODE"))

	if err := checkEnv(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}

	var serviceManager manager.ServiceManager

	if err := serviceManager.Create(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
	if err := serviceManager.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
