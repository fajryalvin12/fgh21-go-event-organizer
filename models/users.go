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
	Username string `json:"username,omitempty" form:"username" binding:"required"`
}

type ChangePassword struct {
	OldPassword     string `json:"oldPassword" form:"oldPassword"`
	NewPassword     string `json:"newPassword" form:"newPassword"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" binding:"eqfield=NewPassword"`
}

func CountUsers(src string) int {
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

	return users
}
func FindUserId(id int) Users {

	db := lib.DB()
	defer db.Close(context.Background())
	sql := `SELECT id, email, password, username from users where id=$1`
	rows := db.QueryRow(
		context.Background(),
		sql,
		id,
	)

	user := Users{}
	rows.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.Username,
	)

	return user
}
func CreateNewUser(data Users) Users {
	db := lib.DB()
	defer db.Close(context.Background())
	data.Password = lib.Encrypt(data.Password)

	sql := `insert into "users" ("email", "username", "password") values ($1, $2, $3) returning "id", "email", "username"`
	row := db.QueryRow(context.Background(), sql, data.Email, data.Username, data.Password)

	var results Users
	row.Scan(
		&results.Id,
		&results.Email,
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

	users, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[Users])

	user := Users{}
	for _, val := range users {
		if val.Email == email {
			user = val
		}
	}
	return user
}
func ChangePass(data ChangePassword, id int) Users {
	db := lib.DB()
	defer db.Close(context.Background())

	updatePass := lib.Encrypt(data.NewPassword)
	sql := `update "users" set "password" = $1 where id=$2`
	fmt.Println(updatePass)

	_, err := db.Exec(context.Background(), sql, updatePass, id)

	if err != nil {
		fmt.Println(err)
		return Users{}
	}

	result := FindUserId(id)

	return result
}
