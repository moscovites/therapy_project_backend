package controllers

import (
	"core/models"
	"core/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)


func CreateAppointment(context *gin.Context) {
	var input models.Appointment
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedAppointment, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"data": savedAppointment})
	
}

func UpdateAppointment(context *gin.Context) {
	var input models.Appointment

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := utils.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existingAppointment, err := models.FindAppointmentById(input.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	
	updatedAppointment, err := existingAppointment.Update(&input)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedAppointment})
}


func GetAllAppointments(context *gin.Context) {
    appointments, err := models.AllAppointments()
    if err != nil {
		log.Println("Error fetching appointments:", err)
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"data": appointments})
}