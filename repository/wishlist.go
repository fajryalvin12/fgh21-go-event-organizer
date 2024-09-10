package repository

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindAllWishlist () []models.Wishlist {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "wishlist"`

	query, _ := db.Query(context.Background(), sql)

	rows, _ := pgx.CollectRows(query, pgx.RowToStructByPos[models.Wishlist])

	return rows
}
func FindAllUsersWishlist (userId int) []models.JoinWishlistEvent {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := 
	`select "w"."id", "e"."title", "e"."date", "e"."location_id", "e"."description" 
	from "wishlist" "w" 
	join "events" "e" 
	on "w"."event_id" = "e"."id" 
	where "w"."user_id" = $1;`

	query, err := db.Query(
		context.Background(),
		sql,
		userId,
	)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := pgx.CollectRows(query, pgx.RowToStructByPos[models.JoinWishlistEvent])
	if err != nil {
		fmt.Println(err)
	}
	return rows
}
func AddNewWishlist (data models.Wishlist) models.Wishlist {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `insert into "wishlist" ("user_id", "event_id") values ($1, $2) returning id, user_id, event_id`

	query := db.QueryRow(context.Background(), sql, data.UserId, data.EventId)

	var result models.Wishlist
	query.Scan(
		&result.Id,
		&result.UserId,
		&result.EventId,
	)

	return result
}
func DeleteTheWishlist (id int) models.Wishlist{
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `delete from "wishlist" where id=$1 returning "id", "user_id", "event_id"`

	query := db.QueryRow(context.Background(), sql, id)

	var result models.Wishlist
	query.Scan(&result.Id, &result.UserId, &result.EventId)

	return result
}