package database

type Account struct {
	Id   int
	Iban string
	Name string
}

type Payment struct {
	Id                   int
	DebtorId             string
	CreditorId           string
	Ammount              float64
	IdempotencyUniqueKey string
	Status               string
}
