package models

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
)

type DetailsTrx struct {
	Id 				int 	`json:"id"`
	TransactionId 	int 	`json:"transactionId" db:"transaction_id"`
	SectionId 		int 	`json:"sectionId" db:"section_id"`
	TicketQty 		int 	`json:"ticketQty" db:"ticket_qty"`
} 

func CreateDetailTransaction(details DetailsTrx) DetailsTrx {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `insert into "transaction_details" ("transaction_id", "section_id", "ticket_qty") values ($1, $2, $3) returning "id", "transaction_id", "section_id", "ticket_qty"`

	row := db.QueryRow(context.Background(), sql, details.TransactionId, details.SectionId, details.TicketQty)

	var results DetailsTrx
	row.Scan(
		&results.Id,
		&results.TransactionId,
		&results.SectionId,
		&results.TicketQty,
	)

	fmt.Println(results)
	return results
}