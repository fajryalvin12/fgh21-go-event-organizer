package routers

import (
	"github.com/fajryalvin12/fgh21-go-event-organizer/controllers"
	"github.com/fajryalvin12/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func ProfileRouter(r *gin.RouterGroup) {
	r.Use(middlewares.AuthMiddleware())
	r.GET("", controllers.DetailUserProfile)
	r.PATCH("", controllers.UpdateProfile)
}
