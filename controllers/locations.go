package controllers

import (
	"strconv"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/repository"
	"github.com/gin-gonic/gin"
)

func ListAllLocations(c *gin.Context) {
	shareLoc := repository.ShowAllLocation()

	lib.HandlerOk(c, "List all locations", nil, shareLoc)
}
func DetailOneLocation (c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	selected, err := repository.GetOneLocationById(id)

	if err != nil {
		lib.HandlerNotFound(c, "Data not found")
		return
	}

	if selected.Id != 0 {
		lib.HandlerOk(c, "Detail event", nil, selected)
	} else {
		lib.HandlerNotFound(c, "Data not found")
	}
}