package controllers

import (
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/fajryalvin12/fgh21-go-event-organizer/repository"
	"github.com/gin-gonic/gin"
)

func DetailUserProfile(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
	profile := repository.FindProfileByUserId(userId)

	lib.HandlerOk(ctx, "Detail user profile", nil, profile)
}
func UpdateProfile(ctx *gin.Context) {
	form := models.JoinProfile{}
	ctx.Bind(&form)
	fmt.Println(form)
	id := ctx.GetInt("userId")

	bagas := repository.EditProfileUsers(models.Users{
		Email:    form.Email,
		Username: form.Username,
	}, id)

	edit := repository.ChangeProfileByUserId(models.Profile{
		FullName:      form.FullName,
		BirthDate:     form.BirthDate,
		Gender:        *form.Gender,
		PhoneNumber:   form.PhoneNumber,
		Profession:    form.Profession,
		NationalityId: form.Nationality,
	}, bagas.Id)

	lib.HandlerOk(ctx, "Success to edit profile", nil, edit)
}
