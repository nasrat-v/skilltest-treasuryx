package database

import (
	_ "github.com/mattn/go-sqlite3"
)

func (x *Database) InsertPayment(payment Payment) (Payment, error) {
	stmt, err := x._sqliteDatabase.Prepare(
		"INSERT INTO payment (id, debtorId, creditorId, ammount, idempotencyUniqueKey, status) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return Payment{}, err
	}
	if _, err := stmt.Exec(
		nil,
		payment.DebtorId,
		payment.CreditorId,
		payment.Ammount,
		payment.IdempotencyUniqueKey,
		payment.Status,
	); err != nil {
		return Payment{}, err
	}
	defer stmt.Close()
	return x.GetPaymentByIdempotency(payment.IdempotencyUniqueKey)
}

func (x *Database) GetPaymentByIdempotency(idempotency string) (Payment, error) {
	payment := Payment{}

	rows, err := x._sqliteDatabase.Query(
		"SELECT id, debtorId, creditorId, ammount, idempotencyUniqueKey, status FROM payment WHERE idempotencyUniqueKey = '" +
			idempotency + "'")
	if err != nil {
		return payment, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&payment.Id,
			&payment.DebtorId,
			&payment.CreditorId,
			&payment.Ammount,
			&payment.IdempotencyUniqueKey,
			&payment.Status,
		); err != nil {
			return payment, err
		}
	}
	return payment, nil
}
