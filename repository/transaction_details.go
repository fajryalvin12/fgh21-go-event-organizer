package repository

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
)

func CreateDetailTransaction(details models.DetailsTrx) models.DetailsTrx {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `insert into "transaction_details" ("transaction_id", "section_id", "ticket_qty") values ($1, $2, $3) returning "id", "transaction_id", "section_id", "ticket_qty"`

	row := db.QueryRow(context.Background(), sql, details.TransactionId, details.SectionId, details.TicketQty)

	var results models.DetailsTrx
	row.Scan(
		&results.Id,
		&results.TransactionId,
		&results.SectionId,
		&results.TicketQty,
	)

	fmt.Println(results)
	return results
}