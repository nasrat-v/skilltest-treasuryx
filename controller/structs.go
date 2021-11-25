package controller

type Payment struct {
	DebtorIban           string  `json:"debtor_iban validate:"required"`
	DebtorName           string  `json:"debtor_name" validate:"required,min=3,max=30"`
	CreditorIban         string  `json:"creditor_iban" validate:"required"`
	CreditorName         string  `json:"creditor_name" validate:"required,min=3,max=30"`
	Ammount              float64 `json:"ammount" validate:"required,numeric"`
	IdempotencyUniqueKey string  `json:"idempotency_unique_key" validate:"required,len=10"`
}
