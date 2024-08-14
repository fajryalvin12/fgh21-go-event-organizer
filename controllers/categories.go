package controllers

import (
	"net/http"
	"strconv"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllCategories(ctx *gin.Context) {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	if limit == 0 {
		limit = 5
	}

	
	cat, count := models.ShowAllCategories(search, limit, page)
	pageInfo := lib.PageInfo{
		TotalData: count ,
		TotalPage: 0,
		Page: page,
		Limit: limit,
		Next: 0,
		Prev: 0,
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List All Categories",
		PageInfo: pageInfo,
		Results: cat,
	})
}
func SelectCategory (ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	selected := models.ShowCategoryById(id)


	if selected.Id != 0 {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Detail Category",
			Results: selected,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Data not found",
		})
	}
}
func AddCategory (ctx *gin.Context) {
	cat := models.Category{}

	ctx.Bind(&cat)
	data := models.CreateNewCategory(cat)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Success created new category",
		Results: data,
	})
}
func UpdateCategory (ctx *gin.Context) {
	param := ctx.Param("id")
    id, _  := strconv.Atoi(param)
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	cat := models.Category{}

	sliceCat, _ := models.ShowAllCategories(search, limit, page)

	ctx.Bind(&cat)
	for _, v := range sliceCat {
		if v.Id == id {
			cat = v
		}
	}

	if cat.Id == 0 {
		ctx.JSON(http.StatusNotFound, lib.Response{
            Success: false,
            Message: "Cannot find the category with id:" + param,
        })
        return
	}

	models.EditCategory(cat.Name, id)
	ctx.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "Success editing user with id: " + param,
        Results: cat,
    })
}
func DeleteCategory (ctx *gin.Context) {
	id, _:= strconv.Atoi(ctx.Param("id"))

	delete := models.RemoveCategory(id)
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