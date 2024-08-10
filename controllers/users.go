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
	users := models.FindAllUsers()
	c.JSON(http.StatusOK, lib.Users{
		Success: true,
		Message: "OK",
		Results: users,
	})
}
func DetailUser (c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	selected := models.FindUserId(id)

	if selected.Id != 0 {
		c.JSON(http.StatusOK, lib.Users{
			Success: true,
			Message: "Detail User",
			Results: selected,
		})
	} else {
		c.JSON(http.StatusNotFound, lib.Users{
			Success: false,
			Message: "Data not found",
		})
	}
} 
func CreateUser (c *gin.Context) {
	user := models.Users{}

	if err := c.ShouldBind(&user)
	err != nil {
		c.JSON(http.StatusNotFound, lib.Users{
			Success: false,
			Message: "data not found",
		})
		return
	}

	createData := models.CreateNewUser(user)
	if createData == user {
		c.JSON(http.StatusNotFound, lib.Users{
			Success: false,
			Message: "user not process",
		})
		return
	}
	c.JSON(http.StatusOK, lib.Users{
		Success: true,
		Message: "Create data success",
		Results: createData,
	})
}
func UpdateUser (c *gin.Context) {
	param := c.Param("id")
    id, _  := strconv.Atoi(param)
    data := models.FindAllUsers()

    user := models.Users{}
    err := c.Bind(&user)
    if err != nil {
        fmt.Println(err)
        return
    }

    result := models.Users{}
    for _, v := range data {
        if v.Id == id {
            result = v
        }
    }

    if result.Id == 0 {
        c.JSON(http.StatusNotFound, lib.Users{
            Success: false,
            Message: "Cannot find the user with id:" + param,
        })
        return
    }

    models.EditTheUser(user.Email, user.Username, user.Password, param)

    c.JSON(http.StatusOK, lib.Users{
        Success: true,
        Message: "Success editing user with id: " + param,
        Results: user,
    })
}
func DeleteUser (ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	selectUser := models.FindUserId(id)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, lib.Users{
			Success: false,
			Message: "Data not found",
		})
		return	
	}

	err = models.RemoveUser(models.Users{}, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Users{
			Success: false,
			Message: "Id not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Users{
		Success: true,
		Message: "Successfully deleted the data!",
		Results: selectUser,
	})	
}