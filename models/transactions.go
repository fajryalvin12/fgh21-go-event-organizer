package models

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Transaction struct {
	Id int 					`json:"id"`
	EventId int 			`json:"eventId" db:"event_id"`
	PaymentMethodId int 	`json:"paymentMethodId" db:"payment_method_id"`
	UserId int 				`json:"userId" db:"user_id"`
}
type DetailTransaction struct {
	TransactionId int 		`json:"transactionId" db:"transaction_id"`
	FullName string 		`json:"fullName" db:"full_name"`
	EventTitle string 		`json:"eventTitle" db:"event_title"`
	LocationId *int 			`json:"locationId" db:"location_id"`
	Date string 			`json:"date"`
	PaymentMethod string 	`json:"paymentMethod" db:"payment_method"`
	TicketSection []string 	`json:"section" db:"ticket_section"`
	Quantity []int 			`json:"ticketQty" db:"quantity"`
}

func CreateNewTransactions(data Transaction) Transaction {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `insert into "transactions" ("event_id", "payment_method_id", "user_id") values ($1, $2, $3) returning "id", "event_id", "payment_method_id", "user_id"`
	row := db.QueryRow(context.Background(), sql, data.EventId, data.PaymentMethodId, data.UserId)

	var results Transaction
	row.Scan(
		&results.Id,
		&results.EventId,
		&results.PaymentMethodId,
		&results.UserId,
	)
	fmt.Println(results)
	return results
}
func ListOfTransactions (id int) DetailTransaction {
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
	GROUP BY t.id, p.full_name, e.title, e.location_id, e."date", pm.name `
	row := db.QueryRow(context.Background(), sql, id )

	var collect DetailTransaction
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
func FindTransactionByUserId (id int) []DetailTransaction {
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

	trx, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[DetailTransaction])

	return trx
}