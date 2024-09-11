package repository

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func CreateProfile(data models.Profile) models.JoinProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	sqlProfile := `insert into "profile" 
	("picture","full_name","birth_date","gender","phone_number","profession","nationality_id","user_id") 
	values 
	($1, $2, $3, $4, $5, $6, $7, $8)`


	_, err := db.Exec(context.Background(), sqlProfile, data.Picture, data.FullName, data.BirthDate, data.Gender, data.PhoneNumber, data.Profession, data.NationalityId, data.UserId)

	if err != nil {
		fmt.Println(err)
	}

	var result models.JoinProfile

	result.Id = data.UserId
	result.FullName = data.FullName
	fmt.Println(result)

	return result
}
func ListAllProfile() []models.JoinProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	joinSql := `select "u"."id", "p"."full_name","u"."username", "u"."email", "p"."gender","p"."phone_number", "p"."profession", "p"."nationality_id", "p"."birth_date"  
	from "users" "u" 
	join "profile" "p"
	on "u"."id" = "p"."user_id" order by asc`

	rows, err := db.Query(
		context.Background(),
		joinSql,
	)
	if err != nil {
		fmt.Println(err)
	}
	data, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[models.JoinProfile])
	return data
}
func FindProfileByUserId(id int) models.JoinProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select "u"."id", "p"."full_name","u"."username", "u"."email", "p"."gender","p"."phone_number", "p"."profession", "p"."nationality_id", "p"."birth_date"  
	from "users" "u" 
	join "profile" "p"
	on "u"."id" = "p"."user_id" where "u"."id" = $1`

	row := db.QueryRow(context.Background(), sql, id)

	var result models.JoinProfile
	row.Scan(
		&result.Id,
		&result.FullName,
		&result.Username,
		&result.Email,
		&result.Gender,
		&result.PhoneNumber,
		&result.Profession,
		&result.Nationality,
		&result.BirthDate,
	)

	return result
}
func ChangeProfileByUserId(data models.Profile, id int) models.JoinProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `update "profile" set ("full_name", "phone_number", "gender", "profession", "nationality_id", "birth_date") = ($1, $2, $3, $4, $5, $6) where "user_id"=$7`

	db.Exec(context.Background(), sql, data.FullName, data.PhoneNumber, data.Gender, data.Profession, data.NationalityId, data.BirthDate, id)
	
	result := FindProfileByUserId(id)
	return result
}
