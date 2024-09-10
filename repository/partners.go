package repository

import (
	"context"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindAllPartners()[]models.Partner {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "partners"`

	query, _ := db.Query(
		context.Background(),
		sql,
	)

	rows, _ := pgx.CollectRows(query, pgx.RowToStructByPos[models.Partner])

	return rows
}