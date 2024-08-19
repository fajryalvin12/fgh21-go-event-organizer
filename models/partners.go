package models

import (
	"context"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Partner struct {
	Id int `json:"id"`
	Images string `json:"images"`
}

func FindAllPartners()[]Partner {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "partners"`

	query, _ := db.Query(
		context.Background(),
		sql,
	)

	rows, _ := pgx.CollectRows(query, pgx.RowToStructByPos[Partner])

	return rows
}