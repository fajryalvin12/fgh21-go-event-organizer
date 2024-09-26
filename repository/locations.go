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
func GetOneLocationById(id int) (models.Locations, error){
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM locations WHERE id=$1`
	query, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return models.Locations{}, err
	}
	row, err := pgx.CollectOneRow(query, pgx.RowToStructByName[models.Locations])
	if err != nil {
		return models.Locations{}, err
	}
	return row, nil
}