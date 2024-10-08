package routers

import (
	"github.com/fajryalvin12/fgh21-go-event-organizer/controllers"
	"github.com/fajryalvin12/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func TransactionsRouter(r *gin.RouterGroup) {
	r.Use(middlewares.AuthMiddleware())
	r.POST("", controllers.CreateTransaction)
	r.GET("", controllers.ListOfTransactionsByUserId)
}