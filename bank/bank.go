package bank

import (
	"encoding/csv"
	"encoding/xml"
	"io"
	"os"
	"skilltest-treasuryx/database"
	"time"
)

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

func CreateXmlFile(document Document) error {
	filename := "payment.xml"
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

func GetBankStatusResponse(id string) (string, error) {
	csvFile, err := os.Open("bank_response.csv")
	if err != nil {
		return "", err
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return "", err
	}
	for _, line := range csvLines {
		status := Status{
			Id:     line[0],
			Status: line[1],
		}
		if status.Id == id {
			return status.Status, nil // return status matching with payment id
		}
	}
	return "", nil
}
