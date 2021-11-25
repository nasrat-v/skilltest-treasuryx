package bank

import (
	"encoding/csv"
	"encoding/xml"
	"io"
	"os"
	"skilltest-treasuryx/src/database"
	"time"
)

// Transform creditor, debtor and payment infos to XML format
func MarshalDocument(creditorAccount database.Account, debtorAccount database.Account, payment database.Payment) Document {
	return Document{
		GrpHdr: GrpHdr{
			MsgId:   payment.IdempotencyUniqueKey,
			CreDtTm: time.Now().String(),
		},
		Cdtr: Cdtr{
			Nm: creditorAccount.Name,
			CdtrAcct: CdtrAcct{
				Id: Id{
					IBAN: creditorAccount.Iban,
				},
			},
		},
		Dbtr: Dbtr{
			Nm: debtorAccount.Name,
			CdtrAcct: CdtrAcct{
				Id: Id{
					IBAN: debtorAccount.Iban,
				},
			},
		},
		Amt: payment.Ammount,
	}
}

// Create XML file for the bank
func CreateXmlFile(id string, document Document) error {
	filename := os.Getenv("BANK_FOLDER") + id + "_payment.xml"
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	xmlWriter := io.Writer(file)
	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	if err := enc.Encode(document); err != nil {
		return err
	}
	return nil
}

// Open XML response file from the bank
// Return status related to payment id
// boolean set to true if file exist, false if not
func GetBankStatusResponse(id string) (string, bool, error) {
	filename := os.Getenv("BANK_FOLDER") + "bank_response.csv"

	if _, err := os.Stat(filename); err != nil { // Return if file doesn't exist
		return "", false, nil
	}

	csvFile, err := os.Open(filename) // Open file if exist
	if err != nil {
		return "", true, err
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return "", true, err
	}
	for _, line := range csvLines {
		status := Status{
			Id:     line[0],
			Status: line[1],
		}
		if status.Id == id {
			return status.Status, true, nil // return status matching with payment id
		}
	}
	return "", true, nil
}
