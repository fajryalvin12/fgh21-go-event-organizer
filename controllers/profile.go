package controllers

import (
	"net/http"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)
type FormUpdateProfile struct {
	FullName 		string 	`form:"fullName`
	Username 		*string `form:"username"`
	Email 			string 	`form:"email"`
	Gender 			int 	`form:"gender"`
	PhoneNumber 	string 	`form:"phoneNumber"`
	Profession		string 	`form:"profession"`
	Nationality		int 	`form:"nationality"`
	BirthDate 		string 	`form:"birthDate"`
}

func DetailUserProfile(ctx *gin.Context) {

	userId := ctx.GetInt("userId")
	profile := models.FindProfileByUserId(userId)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Detail user profile",
		Results: profile,
	})
}
func UpdateProfile (ctx *gin.Context) {
	form := models.JoinProfile{}
	ctx.Bind(&form)
	id := ctx.GetInt("userId")
	
	models.EditProfileUsers(models.Users{
		Username: &form.Username,
		Email: form.Email,
	}, id)

	edit := models.ChangeProfileByUserId(models.Profile{
		FullName: form.FullName,
		BirthDate: form.BirthDate,
		Gender: *form.Gender,
		PhoneNumber: form.PhoneNumber,
		Profession: form.Profession,
		NationalityId: form.Nationality,
	}, id)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Success to edit profile",
		Results: edit,
	})
}