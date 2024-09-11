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
func FindProfileByUserId(id int) (models.JoinProfile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select "u"."id","p"."picture", "p"."full_name","u"."username", "u"."email", "p"."gender","p"."phone_number", "p"."profession", "p"."nationality_id", "p"."birth_date"  
	from "users" "u" 
	join "profile" "p"
	on "u"."id" = "p"."user_id" where "u"."id" = $1`

	query, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return models.JoinProfile{}, err
	}
	
	row, err := pgx.CollectOneRow(query, pgx.RowToStructByName[models.JoinProfile])
	
	if err != nil {
		fmt.Println("sini")
		return models.JoinProfile{}, err
	}

	return row, nil
}
func ChangeProfileByUserId(data models.Profile, id int) models.Profile {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `update "profile" set ("full_name", "phone_number", "gender", "profession", "nationality_id", "birth_date") = ($1, $2, $3, $4, $5, $6) where "user_id"=$7 returning *`

	query, err := db.Query(context.Background(), sql, data.FullName, data.PhoneNumber, data.Gender, data.Profession, data.NationalityId, data.BirthDate, id)

	if err != nil {
		return models.Profile{}
	}

	row, err := pgx.CollectOneRow(query, pgx.RowToStructByName[models.Profile])

	if err != nil {
		return models.Profile{}
	}
	
	return row
}
func UploadProfilePicture(data models.Profile, id int) (models.Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE profile SET picture = $1 WHERE user_id = $2 RETURNING *`

	query, err := db.Query(context.Background(), sql, data.Picture, id)

	if err != nil {
		return models.Profile{}, nil
	}

	row, err := pgx.CollectOneRow(query, pgx.RowToStructByPos[models.Profile])

	if err != nil {
		return models.Profile{}, nil
	}

	return row, nil
}