package lib

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
	connString := "postgresql://postgres:1@103.93.58.89:54322/event_organizer?sslmode=disable"
	conn, _ := pgx.Connect(
		context.Background(),
		connString,
	)

	return conn
}
