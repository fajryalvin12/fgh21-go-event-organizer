package controllers

import (
	"fmt"
	"strconv"

	"github.com/fajryalvin12/fgh21-go-event-organizer/dtos"
	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/fajryalvin12/fgh21-go-event-organizer/repository"
	"github.com/gin-gonic/gin"
)

func ListEventsWithPagination(ctx *gin.Context) {
	search := ctx.Query("search")
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}

	events := repository.FindEventWithPagination(search, limit, page)
	lib.HandlerOk(ctx, "List all events", nil, events)
}
func DetailEvent (ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	selected := repository.FindEventById(id)

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
	fmt.Println(&createdBy)

	result := repository.CreateNewEvent(models.Events{
		Image: newEvent.Image,
		Title: newEvent.Title,
		Date: newEvent.Date,
		Description: newEvent.Description,
		CreatedBy: &createdBy,
	})
	lib.HandlerOk(ctx, "Success create new event!", nil, result)
}
func UpdateEvent (ctx *gin.Context) { 
    id, _  := strconv.Atoi(ctx.Param("id"))
	selected := models.Events{}
	ctx.Bind(&selected)

	updated := repository.EditTheEvent(selected, id)

	if updated.Id == 0 {
		lib.HandlerNotFound(ctx, "Event not found")
		return
	}
	lib.HandlerOk(ctx, "Success edit event", nil, updated)
}
func DeleteEvent (ctx *gin.Context) {
	id, _:= strconv.Atoi(ctx.Param("id"))
	createdBy := ctx.GetInt("userId")

	event := repository.FindEventById(id)

	if createdBy != *event.CreatedBy {
		lib.HandlerBadRequest(ctx, "Cannot deleted event from another user")
		return
	}

	delete := repository.RemoveTheEvent(id)
	if delete.Id == 0 {
		lib.HandlerNotFound(ctx, "Cannot delete the data due to failed request")
		return
	}
	fmt.Println(delete)

	lib.HandlerOk(ctx, "Success deleted the data", nil, delete)
}
func ListPaymentMethods (ctx *gin.Context) {
	payment := repository.FindAllPaymentMethods()
	lib.HandlerOk(ctx, "List all payment methods", nil, payment)
}
func ListAllSectionsByEvent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	sections := repository.FindAllSectionsByEventId(id)
	lib.HandlerOk(ctx, "List all sections", nil, sections)
}
func CreateNewSectionByEventId (ctx *gin.Context) {
	form := dtos.FormSection{}

	fmt.Println(form.EventId)
	err := ctx.Bind(&form)

	if err != nil {
		lib.HandlerBadRequest(ctx, "Please input the data first")
		return
	}

	sections, err := repository.CreateNewSection(models.Section{
		EventId: form.EventId,
		SectionName: form.SectionName,
		Quantity: form.Quantity,
		SectionPrice: form.SectionPrice,
	})

	if err != nil {
		fmt.Println(err)
		lib.HandlerBadRequest(ctx, "data not proper")
		return
	}

	data := repository.FindSectionByEventId(sections.Id)

	lib.HandlerOk(ctx, "Success add new section", nil, data)
}
func ShowEventsByCategory (ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	cat, err := repository.FindEventFromCategory(id)
	if err != nil {
		lib.HandlerNotFound(ctx, "Data not found")
		return
	}
	lib.HandlerOk(ctx, "List The Events from Selected Category", nil, cat)
}