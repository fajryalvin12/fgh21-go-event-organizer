package routers

import (
	"github.com/fajryalvin12/fgh21-go-event-organizer/controllers"
	"github.com/fajryalvin12/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func CategoriesRouter(r *gin.RouterGroup) {
	r.Use(middlewares.AuthMiddleware())
	r.GET("", controllers.ListAllCategories)
	r.GET("/:id", controllers.SelectCategory)
	r.POST("", controllers.AddCategory)
	r.PATCH("/:id", controllers.UpdateCategory)
	r.DELETE("/:id", controllers.DeleteCategory)
}