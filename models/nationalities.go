package models

import (
	"context"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Nationalities struct {
	Id 			int		`json:"id"` 
	Name 		string	`json:"name"`
}

func ShowTheNationalities() []Nationalities {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "nationalities"`

	query, _ := db.Query(
		context.Background(),
		sql,
	)

	rows, _ := pgx.CollectRows(query, pgx.RowToStructByPos[Nationalities])

	return rows
}