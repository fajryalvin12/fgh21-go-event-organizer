package models

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)
type Profile struct {
	Id 				int `json:"id" db:"id"`
	Picture 		*string `json:"picture" form:"picture" db:"picture"`
	FullName 		string `json:"fullName" form:"fullName" db:"full_name"`
	BirthDate 		*string `json:"birthDate" form:"birthDate" db:"birth_date"`
	Gender 			int `json:"gender" form:"gender" db:"gender"`
	PhoneNumber 	*string `json:"phoneNumber" form:"phoneNumber" db:"phone_number"`
	Profession		*string `json:"profession" form:"profession"`
	NationalityId 	*int `json:"nationalityId" form:"nationalityId" db:"nationality_id"`
	UserId 			int `json:"userId" form:"userId" db:"user_id"`
}
type JoinProfile struct {
	Id 				int `json:"id"`
	FullName 		string `json:"fullName" db:"full_name" form:"fullName"`
	Username 		string `json:"username,omitempty" db:"username" form:"userName"`
	Email 			string `json:"email" form:"email"`
	Gender 			*int `json:"gender,omitempty" form:"gender"`
	PhoneNumber 	*string `json:"phoneNumber,omitempty" db:"phone_number" form:"phoneNumber"`
	Profession		*string `json:"profession,omitempty" form:"profession"`
	Nationality		*int `json:"nationality,omitempty" form:"nationality"`
	BirthDate 		*string `json:"birthDate,omitempty" db:"birth_date" form:"birthDate"`
}

func CreateProfile(data Profile) JoinProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	sqlProfile := `insert into "profile" 
	("picture","full_name","birth_date","gender","phone_number","profession", "nationality_id", "user_id") 
	values 
	($1, $2, $3, $4, $5, $6, $7, $8)`

	db.Exec(context.Background(), sqlProfile, data.Picture, data.FullName, data.BirthDate, data.Gender, data.PhoneNumber, data.Profession, data.NationalityId, data.UserId)

	var result JoinProfile

	result.Id = data.UserId
	result.FullName = data.FullName

	return result
}
func ListAllProfile ()[]JoinProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	joinSql := `select "u"."id", "p"."full_name","u"."username", "u"."email", "p"."gender","p"."phone_number", "p"."profession", "p"."nationality_id", "p"."birth_date"  
	from "users" "u" 
	join "profile" "p"
	on "u"."id" = "p"."user_id" order by asc`
		
	rows, err:= db.Query(
		context.Background(),
		joinSql,
		)
	if err != nil {
		fmt.Println(err)
	}
	data, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[JoinProfile])
	return data
}
func FindProfileByUserId(id int) JoinProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select "u"."id", "p"."full_name","u"."username", "u"."email", "p"."gender","p"."phone_number", "p"."profession", "p"."nationality_id", "p"."birth_date"  
	from "users" "u" 
	join "profile" "p"
	on "u"."id" = "p"."user_id" where "u"."id" = $1`

	row := db.QueryRow(context.Background(), sql, id)

	var result JoinProfile
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
func ChangeProfileByUserId (data Profile, id int) JoinProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `update "profile" set ("full_name", "phone_number", "gender", "profession", "nationality_id", "birth_date") = ($1, $2, $3, $4, $5, $6) where "user_id"=$7`

	db.Exec(context.Background(), sql, data.FullName, data.PhoneNumber, data.Gender, data.Profession, data.NationalityId, data.BirthDate ,id)

	result := FindProfileByUserId(id)
	return result
}