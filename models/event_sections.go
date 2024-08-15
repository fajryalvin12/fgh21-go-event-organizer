package models

import (
	"context"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type EventSection struct {
	Id       	int `json:"id"`
	Name     	string `json:"name"`
	Price    	int	`json:"price"`
	Quantity 	int	`json:"quantity"`
	EventId 	int `json:"eventId" db:"event_id"`
}
func FindAllSectionsByEventId(id int) []EventSection {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "event_sections" where "event_id" = $1`

	rows, _ := db.Query(
		context.Background(),
		sql,
		id,
	)

	sections, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[EventSection])
	
	return sections
}