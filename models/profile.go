package models

import (
	"context"
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
)
type JoinProfile struct {
	Id 				int `json:"id"`
	Email 			string `json:"email" form:"email" db:"email"`
	FullName 		string `json:"fullName" form:"fullName" db:"full_name"`
}

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

func CreateProfile(data Profile) (JoinProfile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sqlProfile := `insert into "profile" 
	("picture", 
	"full_name", 
	"birth_date", 
	"gender", 
	"phone_number",
	"profession", 
	"nationality_id", 
	"user_id") 
	values 
	($1, $2, $3, $4, $5, $6, $7, $8) 
	returning 
	"picture", 
	"full_name", 
	"birth_date", 
	"gender", 
	"phone_number", 
	"profession" 
	"nationality_id", 
	"user_id"`

	_, err := db.Exec(
			context.Background(), 
			sqlProfile, 
			data.Picture, 
			data.FullName, 
			data.BirthDate, 
			data.Gender, 
			data.PhoneNumber, 
			data.Profession, 
			data.NationalityId, 
			data.UserId,
		)

	if err != nil {
		fmt.Println(err)
	}

	joinSql := `select "u"."id", "p"."email", "p"."full_name" 
	from "profile" "p" 
	join "users" "u" 
	on "p"."users_id" = "u"."id"`
		
	joinRow:= db.QueryRow(
		context.Background(),
		joinSql,
		)

	var results JoinProfile
	joinRow.Scan(
		&results.Id,
		&results.FullName,
		&results.Email,
	)
	return results, err
}
func FindProfileByUserId(id int) {
	db := lib.DB()
	defer db.Close(context.Background())
}