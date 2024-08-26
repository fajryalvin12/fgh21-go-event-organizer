package routers

import (
	"github.com/fajryalvin12/fgh21-go-event-organizer/controllers"
	"github.com/fajryalvin12/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func WishlistRouter(r *gin.RouterGroup) {
	r.Use(middlewares.AuthMiddleware())
	r.GET("", controllers.ListAllWishlist)
	r.POST("", controllers.CreateWishlist)
	r.DELETE("/:id", controllers.RemoveWishlist)
}