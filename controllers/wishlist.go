package controllers

import (
	"fmt"
	"net/http"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllWishlist (c *gin.Context) {
	id := c.GetInt("userId")
	list := models.FindAllUsersWishlist(id)

	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List of Wishlist",
		Results: list,
	})
}
func CreateWishlist(c *gin.Context) {
	form := models.Wishlist{}

	c.Bind(&form)
	form.UserId = c.GetInt("userId")
	fmt.Println(form.UserId)
	wish := models.AddNewWishlist(form)

	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "New Wishlist has been added",
		Results: wish,
	})
}