package repository

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func CreateNewTransactions(data models.Transaction) models.Transaction {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `insert into "transactions" ("event_id", "payment_method_id", "user_id") values ($1, $2, $3) returning "id", "event_id", "payment_method_id", "user_id"`
	row := db.QueryRow(context.Background(), sql, data.EventId, data.PaymentMethodId, data.UserId)

	var results models.Transaction
	row.Scan(
		&results.Id,
		&results.EventId,
		&results.PaymentMethodId,
		&results.UserId,
	)
	return results
}
func ListOfTransactions (id int) models.DetailTransaction {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT t.id, p.full_name, e.title as "title_event",
	e.location_id, e."date", pm.name as "payment_method",
	array_agg(es.name) as "section", array_agg(td.ticket_qty) "quantity" 
	FROM transactions t 
	JOIN users u ON u.id = t.user_id
	JOIN profile p ON p.user_id = u.id
	JOIN events e ON e.id = t.event_id
	JOIN payment_methods pm ON pm.id = t.payment_method_id
	JOIN transaction_details td ON td.transaction_id = t.id
	JOIN event_sections es ON es.id = td.section_id
	WHERE t.id = $1
	GROUP BY t.id, p.full_name, e.title, e.location_id, e."date", pm.name`
	row := db.QueryRow(context.Background(), sql, id )

	var collect models.DetailTransaction
	row.Scan(
		&collect.TransactionId,
		&collect.FullName,
		&collect.EventTitle,
		&collect.LocationId,
		&collect.Date,
		&collect.PaymentMethod,
		&collect.TicketSection,
		&collect.Quantity,
	)
	fmt.Println(collect)
	
	return collect
}
func FindTransactionByUserId (id int) []models.DetailTransaction {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT t.id, p.full_name, e.title as "title_event",
	e.location_id, e."date", pm.name as "payment_method",
	array_agg(es.name) as "section", array_agg(td.ticket_qty) "quantity" 
	FROM transactions t 
	JOIN users u ON u.id = t.user_id
	JOIN profile p ON p.user_id = u.id
	JOIN events e ON e.id = t.event_id
	JOIN payment_methods pm ON pm.id = t.payment_method_id
	JOIN transaction_details td ON td.transaction_id = t.id
	JOIN event_sections es ON es.id = td.section_id
	WHERE u.id = $1
	GROUP BY t.id, p.full_name, e.title, e.location_id, e."date", pm.name `
	rows, _ := db.Query(context.Background(), sql, id )

	trx, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[models.DetailTransaction])

	return trx
}