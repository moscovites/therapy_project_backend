package controllers

import (
	"core/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePatientProfile(context *gin.Context) {
	var input models.PatientProfile
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patientId := input.UserID
	_, err := models.FindUserById(patientId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
	}

	savedProfile, err := input.Save()

	context.JSON(http.StatusBadRequest, gin.H{"data": savedProfile})
	
}