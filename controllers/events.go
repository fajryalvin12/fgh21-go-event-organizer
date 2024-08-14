package controllers

import (
	"net/http"
	"strconv"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListEvents(ctx *gin.Context) {
	events := models.FindAllEvents()

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List all events",
		Results: events,
	})
}
func DetailEvent (ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	selected := models.FindEventById(id)

	if selected.Id != 0 {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Detail event",
			Results: selected,
		})
	} else {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: false,
			Message: "Data not found",
		})
	}
}
func CreateEvent (ctx *gin.Context) {
	newEvent := models.Events{}
	ctx.Bind(&newEvent)

	createdBy, _ := ctx.Keys["userId"].(int)
	newEvent.CreatedBy = &createdBy

	result := models.CreateNewEvent(newEvent)
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Success create new event!",
		Results: result,
	})
}
func UpdateEvent (ctx *gin.Context) { 
    id, _  := strconv.Atoi(ctx.Param("id"))
	selected := models.Events{}
	ctx.Bind(&selected)

	updated := models.EditTheEvent(selected, id)

	if updated.Id == 0 {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Event not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Success edit event",
		Results: updated,
	})
}
func DeleteEvent (ctx *gin.Context) {
	id, _:= strconv.Atoi(ctx.Param("id"))

	delete := models.RemoveTheEvent(id)
	if delete.Id == 0 {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Cannot delete the data due to failed request",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Success deleted the data",
		Results: delete,
	})
}