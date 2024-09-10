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

func ListAllUsers(c *gin.Context) {
	search := c.Query("search")
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	if page == 0 {
		page = (page - 1) * limit
	}

	users := repository.FindAllUsers(search, limit, page)

	lib.HandlerOk(c, "List all users", nil, users)
}
func DetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	selected := repository.FindUserId(id)

	if selected.Id != 0 {
		lib.HandlerOk(c, "Detail user", nil, selected)
	} else {
		lib.HandlerNotFound(c, "Data not found")
	}
}
func CreateUser(c *gin.Context) {
	user := models.Users{}

	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := repository.CreateNewUser(user)
	lib.HandlerOk(c, "Create data success", nil, data)
}
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.Users{}
	c.Bind(&user)

	data := repository.EditTheUser(user, id)

	if data.Id == 0 {
		lib.HandlerBadRequest(c, "Cannot find the user with this id")
		return
	}

	lib.HandlerOk(c, "Success edit user", nil, data)
}
func DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	selectUser := repository.FindUserId(id)

	if err != nil {
		lib.HandlerNotFound(ctx, "Data not found")
		return
	}

	err = repository.RemoveUser(models.Users{}, id)
	if err != nil {
		lib.HandlerBadRequest(ctx, "Id not found")
		return
	}

	lib.HandlerOk(ctx, "Success remove the user", nil, selectUser)
}
func ChangePassUser(ctx *gin.Context) {
	form := dtos.ChangePassword{}
	userId := ctx.GetInt("userId")
	err := ctx.Bind(&form)
	if err != nil {
		lib.HandlerBadRequest(ctx, "Please input the password")
		return
	}

	if userId <= 0 {
		lib.HandlerNotFound(ctx, "Data not found")
		return
	}
	user := repository.FindUserId(userId)
	if err != nil {
		lib.HandlerNotFound(ctx, "Data not found")
		return
	}

	isVerified := lib.Verify(form.OldPassword, user.Password)

	if !isVerified {
		lib.HandlerBadRequest(ctx, "Please input match password")
		return
	} 
		pass := repository.ChangePass(form, userId)

		lib.HandlerOk(ctx, "Password has been changed", nil, pass)
}
