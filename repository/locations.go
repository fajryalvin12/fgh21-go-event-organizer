package repository

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func ShowAllLocation() []models.Locations {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "locations"`

	query, _ := db.Query(context.Background(), sql)

	rows, err := pgx.CollectRows(query, pgx.RowToStructByPos[models.Locations])
	if err != nil {
		fmt.Println(err)
	}

	return rows
}