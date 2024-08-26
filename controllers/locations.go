package controllers

import (
	"net/http"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllLocations(c *gin.Context) {
	shareLoc := models.ShowAllLocation()
	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Locations",
		Results: shareLoc,
	})
}