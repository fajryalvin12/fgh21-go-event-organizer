package repository

import (
	"context"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func ShowTheNationalities() []models.Nationalities {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "nationalities"`

	query, _ := db.Query(
		context.Background(),
		sql,
	)

	rows, _ := pgx.CollectRows(query, pgx.RowToStructByPos[models.Nationalities])

	return rows
}