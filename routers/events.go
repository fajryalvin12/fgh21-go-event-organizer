package routers

import (
	"github.com/fajryalvin12/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func EventsRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListEvents)
	r.GET("/:id", controllers.DetailEvent)
	r.POST("", controllers.CreateEvent)
	r.PATCH("/:id",controllers.UpdateEvent)
	r.DELETE("/:id",controllers.DeleteEvent)
}