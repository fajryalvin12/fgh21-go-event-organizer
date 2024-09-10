package repository

import (
	"context"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindAllSectionsByEventId(id int) []models.EventSection {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "event_sections" where "event_id" = $1`

	rows, _ := db.Query(
		context.Background(),
		sql,
		id,
	)

	sections, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[models.EventSection])
	
	return sections
}