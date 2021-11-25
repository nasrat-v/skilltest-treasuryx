package database

type Account struct {
	Id   int
	Iban string
	Name string
}

type Payment struct {
	Id                   int
	DebtorId             int
	CreditorId           int
	Ammount              float64
	IdempotencyUniqueKey string
	Status               string
}
