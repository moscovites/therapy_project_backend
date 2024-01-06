package controllers

import (
	"github.com/gin-gonic/gin"
	"core/models"
	"net/http"
)

func CreateWaitlistMember(context *gin.Context) {
	var input models.Waitlist
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	waitlistMember := models.Waitlist{
		FullName: input.FullName,
		Email: input.Email,
		
	}

	savedWaitlistMember, err := waitlistMember.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Member created successfully", "Created member": savedWaitlistMember})
}