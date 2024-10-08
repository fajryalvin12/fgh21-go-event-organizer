package repository

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindEventWithPagination(search string, limit int, page int) []models.EventLocation {
	db := lib.DB()
	defer db.Close(context.Background())

	offset := (page - 1) * limit
	sql := `SELECT e.id, e.image, e.title, e.date, e.description, l.name as location, e.created_by FROM events e
			JOIN locations l ON e.location_id = l.id
			WHERE e.title ILIKE '%' || $1 || '%'
			LIMIT $2
			OFFSET $3;
			`
	rows, _ := db.Query(context.Background(), sql, search, limit, offset)
	
	events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.EventLocation])

	if err != nil {
		fmt.Println(err)
	}

	return events
}
func FindEventById (id int) (models.EventLocation, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT e.id, e.image, e.title, e.date, e.description, l.name as location, e.created_by FROM events e
	JOIN locations l ON e.location_id = l.id
	WHERE e.id=$1
	`
	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return models.EventLocation{}, err
	}

	row, err := pgx.CollectOneRow(query, pgx.RowToStructByName[models.EventLocation])
	
	if err != nil {
		return models.EventLocation{}, err
	}

	return row, err
}
func CreateNewEvent(data models.Events) models.Events {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO events ("image", "title", "date", "description", "created_by") VALUES 
	($1, $2, $3, $4, $5) RETURNING "id", "image", "title", "date", "description", "created_by"`

	query := db.QueryRow(context.Background(), sql, data.Image, data.Title, data.Date, data.Description, data.CreatedBy)

	var result models.Events

	err := query.Scan(
		&result.Id,
		&result.Image,
		&result.Title,
		&result.Date,
		&result.Description,
		&result.CreatedBy,
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
func RemoveTheEvent (id int) models.EventLocation {
	db := lib.DB()
    defer db.Close(context.Background())

	event, _ := FindEventById(id)

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
func CreateNewSection (data models.Section) (models.Section, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT into event_sections("name", "price", "quantity", "event_id") VALUES ($1, $2, $3, $4) RETURNING "id", "name", "price", "quantity", "event_id"`

	query, err := db.Query(context.Background(), sql, data.SectionName, data.SectionPrice, data.Quantity, data.EventId)

	if err != nil {
		return models.Section{}, err
	}

	row, err := pgx.CollectOneRow(query, pgx.RowToStructByPos[models.Section])

	if err != nil {
		return models.Section{}, err
	}

	return row, err
}
func FindEventFromCategory (id int) ([]models.Events, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT ec.id, e.image, e.title, e."date", e.description, e.location_id, e.created_by FROM categories c
			JOIN event_categories ec ON ec.event_id = c.id
			JOIN events e ON ec.event_id = e.id
			WHERE ec.categories_id = $1;`
	
	query, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return []models.Events{}, err
	}

	rows, err := pgx.CollectRows(query, pgx.RowToStructByName[models.Events])
	if err != nil {
		return []models.Events{}, err
	}
	return rows, err
}