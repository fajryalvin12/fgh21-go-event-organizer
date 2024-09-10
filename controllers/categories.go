package controllers

import (
	"strconv"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/fajryalvin12/fgh21-go-event-organizer/repository"
	"github.com/gin-gonic/gin"
)

func ListAllCategories(ctx *gin.Context) {
	cat := repository.ShowAllCategories()
	lib.HandlerOk(ctx, "List All Categories", nil, cat)
}
func SelectCategory (ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	selected := repository.ShowCategoryById(id)

	if selected.Id != 0 {
		lib.HandlerOk(ctx, "Detail Category", nil, selected)
	} else {
		lib.HandlerNotFound(ctx, "Data not found")
	}
}
func AddCategory (ctx *gin.Context) {
	cat := models.Category{}

	ctx.Bind(&cat)
	data := repository.CreateNewCategory(cat)

	lib.HandlerOk(ctx, "Success created new category", nil, data)
}
func UpdateCategory (ctx *gin.Context) {
    id, _  := strconv.Atoi(ctx.Param("id"))
	selected := models.Category{}
	ctx.Bind(&selected)

	update := repository.EditCategory(selected, id)

	if update.Id == 0 {
		lib.HandlerNotFound(ctx, "Data not found")
		return
	}

	lib.HandlerOk(ctx, "Success updated new category", nil, update)
}
func DeleteCategory (ctx *gin.Context) {
	id, _:= strconv.Atoi(ctx.Param("id"))

	delete := repository.RemoveCategory(id)
	if delete.Id == 0 {
		lib.HandlerBadRequest(ctx, "Cannot delete the data due to failed request")
		return
	}
	lib.HandlerOk(ctx, "Success deleted the data", nil, delete)
}