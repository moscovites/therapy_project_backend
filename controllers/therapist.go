package controllers

import (
	"core/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"core/utils"
	"fmt"
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
	fmt.Println("input", input)

	savedProfile, err := input.Save()
	fmt.Println("saved profile", savedProfile)
	context.JSON(http.StatusOK, gin.H{"data": savedProfile})
	
}

func UpdateTherapistProfile(context *gin.Context) {
	var input models.TherapistProfile

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := utils.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existingProfile, err := models.FindTherapistProfileById(input.ID)
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