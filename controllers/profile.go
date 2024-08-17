package controllers

import (
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
func UpdateProfile (ctx *gin.Context) {
	id := ctx.GetInt("userId")
	selected := models.Profile{}
	ctx.Bind(&selected)

	updated := models.ChangeDataProfile(selected, id)
	updated.UserId = id

	if updated.Id == 0 {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Profile not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Success edit profile",
		Results: updated,
	})
}