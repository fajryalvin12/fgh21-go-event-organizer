package controllers

import (
	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllLocations(c *gin.Context) {
	shareLoc := models.ShowAllLocation()

	lib.HandlerOk(c, "List all locations", nil, shareLoc)
}