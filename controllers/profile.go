package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/fajryalvin12/fgh21-go-event-organizer/dtos"
	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/fajryalvin12/fgh21-go-event-organizer/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DetailUserProfile(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
	profile, _ := repository.FindProfileByUserId(userId)

	lib.HandlerOk(ctx, "Detail user profile", nil, profile)
}
func UpdateProfile(ctx *gin.Context) {
	form := dtos.JoinProfile{}
	ctx.Bind(&form)
	fmt.Println(form)
	id := ctx.GetInt("userId")

	user := repository.EditProfileUsers(models.Users{
		Email:    form.Email,
		Username: form.Username,
	}, id)

	
	profile := repository.ChangeProfileByUserId(models.Profile{
		FullName:      form.FullName,
		BirthDate:     &form.BirthDate,
		Gender:        &form.Gender,
		PhoneNumber:   &form.PhoneNumber,
		Profession:    &form.Profession,
		NationalityId: &form.Nationality,
	}, user.Id)
		fmt.Println(profile)
	results, err := repository.FindProfileByUserId(profile.UserId)
	if err != nil {
		fmt.Println(err)
		lib.HandlerNotFound(ctx, "Data not found")
		return
	}

	lib.HandlerOk(ctx, "Success to edit profile", nil, results)
}
func UploadImage (ctx *gin.Context) {
	id := ctx.GetInt("userId")

	maxFile := 2 * 1024 * 1024
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, int64(maxFile))

	file, err := ctx.FormFile("profileImg")
	if err != nil {
		if err.Error() == "http: request body too large" {
			lib.HandlerMaxFile(ctx, "Maximum upload file capacity 2 MB!")
			return
		}
		lib.HandlerBadRequest(ctx, "File doesn't exist!")
	}

	if id == 0 {
		lib.HandlerNotFound(ctx, "User not found!")
		return
	}

	allowExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
	fileExt := strings.ToLower(filepath.Ext(file.Filename))

	if !allowExt[fileExt] {
		lib.HandlerBadRequest(ctx, "The file extension was prohibited!")
		return
	}

	newFile := uuid.New().String() + fileExt
	directory := "./img/profile/"

	if err := ctx.SaveUploadedFile(file, directory + newFile); err != nil {
		lib.HandlerBadRequest(ctx, "Upload failed!")
		return
	}

	fileName := "/img/profile/" + newFile

	delExistingImage, _ := repository.FindProfileByUserId(id)
	if delExistingImage.Picture != nil {
		deleteFile := strings.Split(*delExistingImage.Picture, "8888")[1]
		os.Remove("." + deleteFile)
	}

	profile, _ := repository.UploadProfilePicture(models.Profile{Picture: &fileName}, id)
	fmt.Println(profile)
	lib.HandlerOk(ctx, "Profile picture has been changed", nil, profile)
}
