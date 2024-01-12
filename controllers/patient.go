package controllers

import (
	"core/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"core/utils"
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

	context.JSON(http.StatusOK, gin.H{"data": savedProfile})
	
}

func UpdatePatientProfile(context *gin.Context) {
	var input models.PatientProfile

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := utils.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existingProfile, err := models.FindProfileById(input.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	
	updatedProfile, err := existingProfile.Update(&input)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedProfile})
}