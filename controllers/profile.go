package controllers

import (
	"fmt"
	"net/http"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func DetailUserProfile(ctx *gin.Context) {

	userId := ctx.GetInt("userId")
	profile := models.FindProfileByUserId(userId)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Detail user profile",
		Results: profile,
	})
}
func UpdateProfile(ctx *gin.Context) {
	form := models.JoinProfile{}
	ctx.Bind(&form)
	fmt.Println(form)
	id := ctx.GetInt("userId")

	bagas := models.EditProfileUsers(models.Users{
		Email:    form.Email,
		Username: form.Username,
	}, id)

	edit := models.ChangeProfileByUserId(models.Profile{
		FullName:      form.FullName,
		BirthDate:     form.BirthDate,
		Gender:        *form.Gender,
		PhoneNumber:   form.PhoneNumber,
		Profession:    form.Profession,
		NationalityId: form.Nationality,
	}, bagas.Id)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Success to edit profile",
		Results: edit,
	})
}
