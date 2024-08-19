package controllers

import (
	"net/http"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllPartners(ctx *gin.Context) {
	partners := models.FindAllPartners()
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List of Partners",
		Results: partners,
	})
}