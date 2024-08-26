package lib

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
	connString := "postgresql://postgres:1@172.17.0.2:5432/event_organizer?sslmode=disable"
	conn, _ := pgx.Connect(
		context.Background(),
		connString,
	)

	return conn
}
