package models

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Locations struct {
	Id 			int		`json:"id"` 
	Name 		string	`json:"name"`
	Lat 		string	`json:"lat"`
	Long		string 	`json:"long"`
}

func ShowAllLocation() []Locations {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "locations"`

	query, _ := db.Query(context.Background(), sql)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	rows, err := pgx.CollectRows(query, pgx.RowToStructByPos[Locations])
	if err != nil {
		fmt.Println(err)
	}

	return rows
}