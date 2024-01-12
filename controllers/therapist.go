package controllers

import (
	"core/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTherapistProfile(context *gin.Context) {
	var input models.TherapistProfile
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	therapistId := input.UserID
	_, err := models.FindUserById(therapistId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
	}

	savedProfile, err := input.Save()

	context.JSON(http.StatusBadRequest, gin.H{"data": savedProfile})
	
}