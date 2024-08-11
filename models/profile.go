package models

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)
type Profile struct {
	Id 				int `json:"id" db:"id"`
	Picture 		*string `json:"picture" db:"picture"`
	FullName 		string `json:"fullName" db:"full_name"`
	BirthDate 		*string `json:"birthDate" db:"birth_date"`
	Gender 			int `json:"gender" db:"gender"`
	PhoneNumber 	*string `json:"phoneNumber" db:"phone_number"`
	Profession		*string `json:"profession" db:"profession"`
	NationalityId 	*int `json:"nationalityId" db:"nationality_id"`
	UserId 			int `json:"userId" db:"user_id"`
}

type JoinProfile struct {
	Id 				int `json:"id"`
	FullName 		string `json:"fullName"`
	Email 			string `json:"email"`
	Gender 			int `json:"gender,omitempty"`
	PhoneNumber 	*string `json:"phoneNumber,omitempty"`
	Profession		*string `json:"profession,omitempty"`
	BirthDate 		*string `json:"birthDate,omitempty"`
	Nationality		int `json:"nationality,omitempty"`
}

func CreateProfile(data Profile) JoinProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	sqlProfile := `insert into "profile" 
	("picture","full_name","birth_date","gender","phone_number","profession", "nationality_id", "user_id") 
	values 
	($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := db.Exec(context.Background(), sqlProfile, data.Picture, data.FullName, data.BirthDate, data.Gender, data.PhoneNumber, data.Profession, data.NationalityId, data.UserId)
	
	if err != nil {
		fmt.Println(err)
	}

	result := JoinProfile{}

	result.Id = data.UserId
	result.FullName = data.FullName

	return result
}
func ListAllProfile ()[]JoinProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	joinSql := `select "u"."id", "u"."email", "p"."full_name", "u"."username", "p"."gender", "p"."phone_number","p"."profession", "p"."nationality_id", "p"."birth_date"  
	from "users" "u" 
	join "profile" "p"
	on "u"."id" = "p"."user_id"`
		
	rows, _:= db.Query(
		context.Background(),
		joinSql,
		)
	
	events, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[JoinProfile])
	return events
}
func FindProfileByUserId(id int) JoinProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	var result JoinProfile
	for _, v := range ListAllProfile() {
		if v.Id == id {
			result = v
		}
	}

	return result
}