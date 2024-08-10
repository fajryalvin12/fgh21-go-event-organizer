package lib

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
	connString := "postgresql://postgres:af8d878454e34aa38ff895418dc4d148@localhost:5432/event_organizer?sslmode=disable"
	conn, _ := pgx.Connect(
		context.Background(),
		connString,
	)

	return conn
}