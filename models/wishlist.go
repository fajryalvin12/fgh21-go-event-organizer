package models

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Wishlist struct {
	Id 		int 	`json:"id"`
	UserId 	int		`json:"userId" form:"userId" db:"user_id"`
	EventId int 	`json:"eventId" form:"eventId" db:"event_id"`
}
type JoinWishlistEvent struct {
	Id 			int 	`json:"id"`
	Title 		string	`json:"title"`
	Date 		string	`json:"date"`
	Location 	*int	`json:"location"`
	Description string 	`json:"description"`
}

func FindAllUsersWishlist (userId int) []JoinWishlistEvent {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := 
	`select "w"."id", "e"."title", "e"."date", "e"."location_id", "e"."description" 
	from "wishlist" "w" 
	join "events" "e" 
	on "w"."event_id" = "e"."id" 
	where "w"."id" = $1;`

	query, err := db.Query(
		context.Background(),
		sql,
		userId,
	)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := pgx.CollectRows(query, pgx.RowToStructByPos[JoinWishlistEvent])
	if err != nil {
		fmt.Println(err)
	}
	return rows
}
func AddNewWishlist (data Wishlist) Wishlist {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `insert into "wishlist" ("user_id", "event_id") values ($1, $2) returning id, user_id, event_id`

	query := db.QueryRow(context.Background(), sql, data.UserId, data.EventId)

	var result Wishlist
	err := query.Scan(
		&result.Id,
		&result.UserId,
		&result.EventId,
	)
	if err != nil {
		fmt.Println(err)
	}

	return result
}