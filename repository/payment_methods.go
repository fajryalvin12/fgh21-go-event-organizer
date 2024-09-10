package repository

import (
	"context"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindAllPaymentMethods() []models.PaymentMethods {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "payment_methods"`

	rows, _ := db.Query(
		context.Background(),
		sql,
	)

	payment, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[models.PaymentMethods])

	return payment
}