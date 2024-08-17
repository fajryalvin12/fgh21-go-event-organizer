package controllers

import (
	"net/http"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllNationalities(ctx *gin.Context) {
	nations := models.ShowTheNationalities()

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Nationalities",
		Results: nations,
	})
}