package models

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Events struct {
	Id          int    `json:"id" db:"id"`
	Image       string `json:"image" form:"image" db:"image"`
	Title       string `json:"title" form:"title" db:"title"`
	Date        string `json:"date" form:"date" db:"date"`
	Description string `json:"description" form:"description" db:"description"`
	LocationId  *int   `json:"location_id" form:"location_id" db:"location_id"`
	CreatedBy   *int   `json:"created_by" form:"created_by" db:"created_by"`
}

func FindAllEvents() []Events {
	db := lib.DB()
	defer db.Close(context.Background())

	showSql := `select * from "events" order by "id" asc`

	rows, _ := db.Query(
		context.Background(),
		showSql,
	)
	
	events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Events])

	if err != nil {
		fmt.Println(err)
	}

	return events
}
func FindEventById (id int) Events {
	db := lib.DB()
	defer db.Close(context.Background())

	event := FindAllEvents()
	userEvent := Events{}
	for _, v := range event {
		if v.Id == id {
			userEvent = v
		}
	}
	return userEvent
}
func CreateNewEvent(data Events) Events {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `insert into "events" ("image", "title", "date", "description") values 
	($1, $2, $3, $4)`
	db.Exec(context.Background(), sql, data.Image, data.Title, data.Date, data.Description)

	return data
}
func EditTheEvent(data Events, id int) Events {
	db := lib.DB()
    defer db.Close(context.Background())

	sql := `update "events" set ("image", "title", "date", "description") = ($1, $2, $3, $4) where "id"=$5`

	db.Exec(context.Background(), sql, data.Image, data.Title, data.Date, data.Description, id)

	data.Id = id
	return data
}
func RemoveTheEvent (id int) Events {
	db := lib.DB()
    defer db.Close(context.Background())

	event := FindEventById(id)

	sql := `delete from "events" where "id" = $1`
	db.Exec(context.Background(), sql, id)

	return event
}