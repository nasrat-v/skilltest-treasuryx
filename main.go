package main

import (
	"fmt"
	"os"
	"skilltest-treasuryx/manager"

	"github.com/joho/godotenv"
)

func checkEnv() error {
	/*x._brokerApiURL = os.Getenv("BROKER_API_URL")
	x._triggersDirname = os.Getenv("TRIGGERS_DIRNAME")
	x._eventsDirname = os.Getenv("EVENTS_DIRNAME")
	x._recordingsDirname = os.Getenv("RECORDINGS_DIRNAME")

	if x._brokerApiURL == "" {
		return errors.New("error: No BrokerAPI URL provided")
	}
	if x._triggersDirname == "" {
		return errors.New("error: No Triggers directory name provided")
	}
	if x._eventsDirname == "" {
		return errors.New("error: No Events directory name provided")
	}
	if x._recordingsDirname == "" {
		return errors.New("error: No Recordings directory name provided")
	}*/
	return nil
}

func main() {
	godotenv.Load()

	if err := checkEnv(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-2)
	}

	var serviceManager manager.ServiceManager

	if err := serviceManager.Create(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}
	if err := serviceManager.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}
}
