package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerOk (c *gin.Context, message string, pageInfo any, data any) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		PageInfo: pageInfo,
		Results: data,
	})
}

func HandlerUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, Response{
		Success: false,
		Message: message,
	})
}

func HandlerNotFound (c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: message,
	})
}

func HandlerBadRequest (c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Message: message,
	})
}

func HandlerMaxFile (c *gin.Context, message string) {
	c.JSON(http.StatusRequestEntityTooLarge, Response{
		Success: false,
		Message: message,
	})
}