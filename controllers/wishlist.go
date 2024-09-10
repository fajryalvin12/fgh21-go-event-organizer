package controllers

import (
	"strconv"

	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllWishlist (c *gin.Context) {
	id := c.GetInt("userId")
	list := models.FindAllUsersWishlist(id)

	lib.HandlerOk(c, "Details of Wishlist", nil, list)
}
func CreateWishlist(c *gin.Context) {
	form := models.Wishlist{}

	c.Bind(&form)
	form.UserId = c.GetInt("userId")
	wish := models.AddNewWishlist(form)

	lib.HandlerOk(c, "New wishlist has been added", nil, wish)
}
func RemoveWishlist (c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	del := models.DeleteTheWishlist(id)

	lib.HandlerOk(c, "The wishlist has been removed", nil, del)
}