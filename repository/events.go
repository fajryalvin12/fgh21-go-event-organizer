package repository

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/dtos"
	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindAllEvents() []models.Events {
	db := lib.DB()
	defer db.Close(context.Background())

	showSql := `select * from "events" order by "id" asc`

	rows, _ := db.Query(
		context.Background(),
		showSql,
	)
	
	events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Events])

	if err != nil {
		fmt.Println(err)
	}

	return events
}
func FindEventById (id int) models.Events {
	db := lib.DB()
	defer db.Close(context.Background())

	event := FindAllEvents()
	userEvent := models.Events{}
	for _, v := range event {
		if v.Id == id {
			userEvent = v
		}
	}
	return userEvent
}
func CreateNewEvent(data dtos.Events) models.Events {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO events ("image", "title", "date", "description") VALUES 
	($1, $2, $3, $4) RETURNING id, "image", "title", "date", "description"`

	query := db.QueryRow(context.Background(), sql, data.Image, data.Title, data.Date, data.Description)

	var result models.Events

	err := query.Scan(
		&result.Id,
		&result.Image,
		&result.Title,
		&result.Date,
		&result.Description,
	)
	if err != nil {
		fmt.Println(err)
	}

	return result
}
func EditTheEvent(data models.Events, id int) models.Events {
	db := lib.DB()
    defer db.Close(context.Background())

	sql := `update "events" set ("image", "title", "date", "description") = ($1, $2, $3, $4) where "id"=$5`

	db.Exec(context.Background(), sql, data.Image, data.Title, data.Date, data.Description, id)

	data.Id = id
	return data
}
func RemoveTheEvent (id int) models.Events {
	db := lib.DB()
    defer db.Close(context.Background())

	event := FindEventById(id)

	sql := `delete from "events" where "id" = $1`
	db.Exec(context.Background(), sql, id)

	return event
}
func FindSectionByEventId (eventId int) models.Section {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "event_sections" where "event_id" = $1`
	rows, _ := db.Query(
		context.Background(),
		sql,
		eventId,
	)

	sections, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Section])
	if err != nil {
		fmt.Println(err)
	}

	section := models.Section{}
	for _, val := range sections {
		if val.EventId == eventId{
			section = val
		}
	}
	return section
}