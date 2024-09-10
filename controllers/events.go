package controllers

import (
	"strconv"

	"github.com/fajryalvin12/fgh21-go-event-organizer/dtos"
	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListEvents(ctx *gin.Context) {
	events := models.FindAllEvents()
	lib.HandlerOk(ctx, "List all events", nil, events)
}
func DetailEvent (ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	selected := models.FindEventById(id)

	if selected.Id != 0 {
		lib.HandlerOk(ctx, "Detail event", nil, selected)
	} else {
		lib.HandlerNotFound(ctx, "Data not found")
	}
}
func CreateEvent (ctx *gin.Context) {
	newEvent := dtos.Events{}
	ctx.Bind(&newEvent)

	createdBy := ctx.GetInt("userId")
	newEvent.CreatedBy = &createdBy

	result := models.CreateNewEvent(newEvent)
	lib.HandlerOk(ctx, "Success create new event!", nil, result)
}
func UpdateEvent (ctx *gin.Context) { 
    id, _  := strconv.Atoi(ctx.Param("id"))
	selected := models.Events{}
	ctx.Bind(&selected)

	updated := models.EditTheEvent(selected, id)

	if updated.Id == 0 {
		lib.HandlerNotFound(ctx, "Event not found")
		return
	}
	lib.HandlerOk(ctx, "Success edit event", nil, updated)
}
func DeleteEvent (ctx *gin.Context) {
	id, _:= strconv.Atoi(ctx.Param("id"))

	delete := models.RemoveTheEvent(id)
	if delete.Id == 0 {
		lib.HandlerNotFound(ctx, "Cannot delete the data due to failed request")
		return
	}
	lib.HandlerOk(ctx, "Success deleted the data", nil, delete)
}
func ListAllSectionsByEvent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	sections := models.FindAllSectionsByEventId(id)
	lib.HandlerOk(ctx, "List all sections", nil, sections)
}
func ListPaymentMethods (ctx *gin.Context) {
	payment := models.FindAllPaymentMethods()
	lib.HandlerOk(ctx, "List all payment methods", nil, payment)
}