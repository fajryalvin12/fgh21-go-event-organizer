package controllers

import (
	"fmt"
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
}

func AuthLogin(ctx *gin.Context) {
	var user models.Users
	ctx.Bind(&user)

	found := models.FindUserEmail(user.Email)

	if found == (models.Users{}) {
		ctx.JSON(http.StatusUnauthorized, lib.Users{
			Success: false,
			Message: "wrong email or password",
		})
		return
	}

	isVerified := lib.Verify(user.Password, found.Password)

	if isVerified {
		JWToken := lib.GenerateUserIdToken(found.Id)
		ctx.JSON(http.StatusOK, lib.Users{
			Success: true,
			Message: "Login Success!",
			Results: Token{
				JWToken,
			},
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, lib.Users{
			Success: false,
			Message: "Wrong email or password",
		})
	}
}
func AuthRegister (ctx *gin.Context) {
	form := FormRegister{}
	user := models.Users{}
	profile := models.Profile{}

	err:= ctx.Bind(&form)

	if err != nil {
		fmt.Println(nil)
	}

	form.FullName = profile.FullName
	form.Email = user.Email
	form.Password = user.Password
	createUser := models.CreateNewUser(user)

	userId := createUser.Id
	profile.UserId = userId

	createProfile, _ := models.CreateProfile(profile)

	ctx.JSON(http.StatusOK, lib.Users{
		Success: true,
		Message: "Register Success",
		Results: createProfile,
	})
}