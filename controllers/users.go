package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllUsers(c *gin.Context) {	
	search := c.Query("search")
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	if page == 0 {
		page = (page - 1) * limit
	}

	users := models.FindAllUsers(search, limit, page)

	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "OK",
		Results: users,
	})
}
func DetailUser (c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	selected := models.FindUserId(id)

	if selected.Id != 0 {
		c.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Detail User",
			Results: selected,
		})
	} else {
		c.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "Data not found",
		})
	}
} 
func CreateUser (c *gin.Context) {
	user := models.Users{}

	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := models.CreateNewUser(user)
	fmt.Println(data)

	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Create data success",
		Results: data,
	})
}
func UpdateUser (c *gin.Context) {
    id, _  := strconv.Atoi(c.Param("id"))
    user := models.Users{}
    c.Bind(&user)

	data := models.EditTheUser(user, id)

    if data.Id == 0 {
        c.JSON(http.StatusNotFound, lib.Response{
            Success: false,
            Message: "Cannot find the user with this id",
        })
        return
    }

    c.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "Success editing user",
        Results: data,
    })
}
func DeleteUser (ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	selectUser := models.FindUserId(id)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, lib.Response{
			Success: false,
			Message: "Data not found",
		})
		return	
	}

	err = models.RemoveUser(models.Users{}, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "Id not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Successfully deleted the data!",
		Results: selectUser,
	})	
}