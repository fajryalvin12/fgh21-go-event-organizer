package controllers

import (
	"net/http"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

type Token struct {
	JWToken string `json:"token"`
}
type FormRegister struct {
	FullName		string `form:"fullName"`
	Email			string `form:"email"`
	Password		string `form:"password"`
	ConfirmPassword	string `form:"confirmPassword" binding:"eqfield=Password"`
	Username 		string `form:"username"`
}
type FormLogin struct {
	Email			string `form:"email"`
	Password		string `form:"password"`
}

func AuthLogin(ctx *gin.Context) {
	var user FormLogin
	ctx.Bind(&user)

	found := models.FindUserEmail(user.Email)

	if found == (models.Users{}) {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "wrong email or password",
		})
		return
	}

	isVerified := lib.Verify(user.Password, found.Password)

	if isVerified {
		JWToken := lib.GenerateUserIdToken(found.Id)
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Login Success!",
			Results: Token{
				JWToken,
			},
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Wrong email or password",
		})
	}
}
func AuthRegister (ctx *gin.Context) {
	form := FormRegister{}
	var user models.Users
	var profile models.Profile

	ctx.Bind(&form)

	user.Email = form.Email
	user.Password = form.Password
	profile.FullName = form.FullName
	createUser:= models.CreateNewUser(user)

	userId := createUser.Id
	profile.UserId = userId

	createProfile := models.CreateProfile(profile)
	createProfile.Email = form.Email
	createProfile.FullName = form.FullName 

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Register Success",
		Results: createProfile,
	})
}