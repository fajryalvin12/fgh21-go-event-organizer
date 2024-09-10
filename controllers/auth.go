package controllers

import (
	"fmt"

	"github.com/fajryalvin12/fgh21-go-event-organizer/dtos"
	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

type Token struct {
	JWToken string `json:"token"`
}

func AuthLogin(ctx *gin.Context) {
	var user dtos.FormLogin
	ctx.Bind(&user)

	found := models.FindUserEmail(user.Email)

	if found == (models.Users{}) {
		lib.HandlerUnauthorized(ctx, "Wrong Email or Password!")
		return
	}

	isVerified := lib.Verify(user.Password, found.Password)

	if isVerified {
		JWToken := lib.GenerateUserIdToken(found.Id)
		lib.HandlerOk(ctx, "Login Success!", nil, Token{JWToken})
	} else {
		lib.HandlerUnauthorized(ctx, "Wrong Email or Password!")
	}
}
func AuthRegister (ctx *gin.Context) {
	form := dtos.FormRegister{}
	var user = models.Users{}
	var profile = models.Profile{}

	err := ctx.Bind(&form)
	if err != nil {
		fmt.Println(err)
	}

	user.Email = form.Email
	user.Password = form.Password
	profile.FullName = form.FullName

	createUser:= models.CreateNewUser(user)

	userId := createUser.Id
	profile.UserId = userId

	createProfile := models.CreateProfile(profile)
	createProfile.Email = form.Email
	createProfile.FullName = form.FullName
	
	lib.HandlerOk(ctx, "Register Success", nil, createProfile)
}