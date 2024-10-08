package routers

import (
	"github.com/fajryalvin12/fgh21-go-event-organizer/controllers"
	"github.com/fajryalvin12/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func EventsRouter(r *gin.RouterGroup) {
	r.GET("", controllers.ListEventsWithPagination)
	r.GET("/:id", controllers.DetailEvent)
	r.POST("", middlewares.AuthMiddleware(), controllers.CreateEvent)
	r.PATCH("/:id", middlewares.AuthMiddleware(), controllers.UpdateEvent)
	r.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteEvent)
	r.GET("/payment_method", controllers.ListPaymentMethods)
	r.GET("/section/:id", controllers.ListAllSectionsByEvent)
	r.POST("/section", controllers.CreateNewSectionByEventId)
	r.GET("/category/:id", controllers.ShowEventsByCategory)
}