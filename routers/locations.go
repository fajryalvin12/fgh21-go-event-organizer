package routers

import (
	"github.com/fajryalvin12/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func LocationsRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListAllLocations)
}