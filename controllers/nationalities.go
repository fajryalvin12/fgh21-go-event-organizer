package controllers

import (
	"github.com/fajryalvin12/fgh21-go-event-organizer/lib"
	"github.com/fajryalvin12/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllNationalities(ctx *gin.Context) {
	nations := models.ShowTheNationalities()
	lib.HandlerOk(ctx, "List All Nationalities", nil, nations)
}