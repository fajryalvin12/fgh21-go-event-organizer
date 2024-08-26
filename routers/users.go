package routers

import (
	"github.com/fajryalvin12/fgh21-go-event-organizer/controllers"
	"github.com/fajryalvin12/fgh21-go-event-organizer/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("", controllers.ListAllUsers)
	routerGroup.GET("/:id", controllers.DetailUser)
	routerGroup.POST("", controllers.CreateUser)
	routerGroup.PATCH("/:id", controllers.UpdateUser)
	routerGroup.DELETE("/:id", controllers.DeleteUser)
	routerGroup.PATCH("/change-password", controllers.ChangePassUser)
}
