package controllers

import (
	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllPartners(ctx *gin.Context) {
	partners := models.FindAllPartners()
	lib.HandlerOk(ctx, "List of partners", nil, partners)
}