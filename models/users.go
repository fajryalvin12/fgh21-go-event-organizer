package models

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"-" form:"password" binding:"required,min=8"`
}

func FindAllUsers() []Users {
	db := lib.DB()
	defer db.Close(context.Background())
	sql := `select * from "users" order by "id" asc `
	rows, _ := db.Query(
		context.Background(),
	 	sql,
	)

	users, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[Users])

	return users
}
func FindUserId(id int) Users {

	db := lib.DB()
	defer db.Close(context.Background())
	sql := `select * from "users" where "id" = $1`
	rows, err := db.Query(
		context.Background(),
	 	sql,
		id,
	)
	if err != nil {
		fmt.Println(err)
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Users])

	if err != nil {
		fmt.Println(err)
	}

	user := Users{}
	for _, val := range users {
		if val.Id == id{
			user = val
		}
	}
	return user
}
func CreateNewUser(data Users) Users {
	db := lib.DB()
	defer db.Close(context.Background())
	fmt.Println(data)
	data.Password = lib.Encrypt(data.Password)
	sql := `insert into "users" ("email", "password", "username") values ($1, $2, $3) returning "id", "email", "password", "username"`

	row := db.QueryRow(
		context.Background(), sql, data.Email, data.Password, data.Username)

	var results Users
	row.Scan(
		&results.Id,
		&results.Email,
		&results.Password,
		&results.Username,
	)
	fmt.Println(results)

	return results
}
func EditTheUser(email string, username string, password string, id string) {
    db := lib.DB()
    defer db.Close(context.Background())

    dataSql := `update "users" set (email , username, password) = ($1, $2, $3) where id=$4`

    db.Exec(context.Background(), dataSql, email, username, password, id)
}
func RemoveUser(data Users, id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	delete, err := db.Exec(
		context.Background(),
		`delete from "users" where id=$1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to delete data")
	}
	if delete.RowsAffected() == 0 {
		return fmt.Errorf("data not found")
	} 

	return nil
}
func FindUserEmail(email string) Users {

	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(
		context.Background(),
	 	`select * from "users" where "email" = $1`,
		email,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Users])

	if err != nil {
		fmt.Println(err)
	}

	user := Users{}
	for _, val := range users {
		if val.Email == email{
			user = val
		}
	}
	return user
}