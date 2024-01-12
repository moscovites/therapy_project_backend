package controllers

import (
	"core/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"core/utils"
	"core/database"
	"fmt"
)

func Register(context *gin.Context) {
	var input models.User
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !utils.IsEmailUnique(input.Email) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Sorry, email already exists"})
		return
	}

	verificationCode := utils.GenerateVerificationCode()

	user := models.User{
		Email: input.Email,
		
		VerificationCode: verificationCode,
		Password: input.Password,
	}
	savedUser, err := user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	utils.SendVerificationEmail(user.Email, user.VerificationCode)
		
	context.JSON(http.StatusCreated, gin.H{"message": "A verification code has been sent to the procided email address.", "created user": savedUser})
}

func VerifyEmail(context *gin.Context) {
	verificationCode := context.Param("code")
	verifiedUser, err := utils.FindUserByVerificationCode(verificationCode)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verifiedUser.IsVerified = true
	verifiedUser.VerificationCode = ""

	database.Database.Save(verifiedUser)
	context.JSON(http.StatusOK, gin.H{"message": "Email verification successful.", "verifiedUser": verifiedUser})	
}

func Login(context *gin.Context) {
	var input models.AuthenticationInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := models.FindUserByEmail(input.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.ValidatePassword(input.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwt, err := utils.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	
	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

func PasswordReset(context *gin.Context) {
	var input models.User
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordResetCode := utils.GeneratePasswordResetCode()

	userEmail := input.Email
	fmt.Println("email entered", userEmail)


	user, err := utils.FindUserByEmail(userEmail)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ResetPasswordCode = passwordResetCode
	database.Database.Save(user)
	utils.SendPasswordResetCode(userEmail, passwordResetCode)
	context.JSON(http.StatusOK, gin.H{"message": "Password reset code sent to your email"})	  
}

func PasswordResetConfirm(context *gin.Context) {
	passwordResetCode := context.Param("password-reset-code")
	existingUser, err := utils.FindUserByPasswordResetCode(passwordResetCode)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// existingUser.ResetPasswordCode = ""
	// database.Database.Save(existingUser)
	context.JSON(http.StatusOK, gin.H{"message": "reset code valid", "user": existingUser})

}

func CreateNewPassword(context *gin.Context) {
	var input models.User
	passwordResetCode := context.Param("password-reset-code")
	existingUser, err := utils.FindUserByPasswordResetCode(passwordResetCode)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPassword := input.Password
	updatedPassword, err := existingUser.UpdatePassword(newPassword)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"updated password": updatedPassword})


}

