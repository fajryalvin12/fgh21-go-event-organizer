package controllers

import (
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
	c.Bind(&user)

	createData := models.CreateNewUser(user)
	c.JSON(http.StatusOK, lib.Users{
		Success: true,
		Message: "Create data success",
		Results: createData,
	})
}
func UpdateUser (c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	selected := models.Users{}

	c.Bind(&selected)

	editData := models.EditTheUser(selected, id)

	if editData.Id != 0 {
		c.JSON(http.StatusOK, lib.Users{
			Success: true,
			Message: "Data was updated",
			Results: editData,
		})
	} else {
		c.JSON(http.StatusNotFound, lib.Users{
			Success: false,
			Message: "Id not found",
		})
	}
}
func DeleteUser (c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data := models.RemoveUser(id)

	if data.Id != 0 {
		c.JSON(http.StatusOK, lib.Users {
			Success: true,
			Message: "Success deleted data",
			Results: data,
		})
	} else {
		c.JSON(http.StatusNotFound, lib.Users {
			Success: false,
			Message: "Success deleted data",
		})
	}
}