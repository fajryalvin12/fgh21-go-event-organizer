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
	Password string `json:"-" form:"password" binding:"required,min=8"`
	Username *string `json:"username,omitempty" form:"username" binding:"required"`
}
func CountUsers (src string) int {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT COUNT(*) FROM users WHERE "email" ILIKE '%' || $1 || '%'`
	row := db.QueryRow(context.Background(), sql, src)

	var num int
	row.Scan(&num)
	return num
}
func FindAllUsers(search string, limit int, page int) []Users {
	db := lib.DB()
	defer db.Close(context.Background())
	offset := (page - 1) * limit
	sql := `select * from "users" where "email" ilike '%' || $1 || '%' limit $2 offset $3`
	rows, _ := db.Query(
		context.Background(),
	 	sql,
		search,
		limit,
		offset,
	)

	users, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[Users])
	fmt.Println(users)

	return users
}
func FindUserId(id int) Users {

	db := lib.DB()
	defer db.Close(context.Background())
	sql := `select * from "users" where "id" = $1`
	rows, _ := db.Query(
		context.Background(),
	 	sql,
		id,
	)


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
	data.Password = lib.Encrypt(data.Password)

	sql := `insert into "users" ("email", "password", "username") values ($1, $2, $3) returning "id", "email", "password", "username"`
	row := db.QueryRow(context.Background(), sql, data.Email, data.Password, data.Username)

	var results Users
	row.Scan(
		&results.Id,
		&results.Email,
		&results.Password,
		&results.Username,
	)
	return results
}
func EditTheUser(data Users, id int) Users {
    db := lib.DB()
    defer db.Close(context.Background())

    dataSql := `update "users" set (email , username, password) = ($1, $2, $3) where id=$4 returning "id", "email", "username", "password"`

    query := db.QueryRow(context.Background(), dataSql, data.Email, data.Password, data.Username, id)

	var result Users 
	query.Scan(
		&result.Id,
		&result.Email,
		&result.Password,
		&result.Username,
	)

	return result
}
func EditProfileUsers(data Users, id int) Users {
    db := lib.DB()
    defer db.Close(context.Background())

    dataSql := `update "users" set (email , username) = ($1, $2) where id=$3 returning "id", "email", "username"`

    query := db.QueryRow(context.Background(), dataSql, data.Email, data.Username, id)

	var result Users 
	query.Scan(
		&result.Id,
		&result.Email,
		&result.Username,
	)

	return result
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