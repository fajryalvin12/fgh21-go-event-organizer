package models

import (
	"context"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type PaymentMethods struct {
	Id   int
	Name string
}

func FindAllPaymentMethods() []PaymentMethods {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "payment_methods"`

	rows, _ := db.Query(
		context.Background(),
		sql,
	)

	payment, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[PaymentMethods])

	return payment
}