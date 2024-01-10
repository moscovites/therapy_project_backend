package utils 

import (

	"core/models"
	"core/database"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
	"log"
)

func IsEmailAvailable (email string) bool {
	var existingUsers models.User
	result := database.Database.Where("email = ?", email).First(&existingUsers)
	return 	result.Error == gorm.ErrRecordNotFound
}

func GeneratePasswordResetCode () string {
	code := uuid.New().String()[:8]
	return code
}

func FindUserByPasswordResetCode (code string) (*models.User, error) {
	var user models.User
	result := database.Database.Where("reset_password_code = ?", code).First(&user)
	
	return &user, result.Error
}

func FindUserByEmail (email string) (*models.User, error) {
	var user models.User
	result := database.Database.Where("email = ?", email).First(&user)
	
	return &user, result.Error
}

func SendPasswordResetCode (email, passwordResetCode string) {
	senderEmail := "stakshare.com@gmail.com"
	senderPassword := "ufbzqxrvvuwrjbuh"

	recipientEmail := email

	message := gomail.NewMessage()
	message.SetHeader("From", senderEmail)
	message.SetHeader("To", recipientEmail)
	message.SetHeader("Subject", "Password reset, Stakshare.")
	message.SetBody("text/plain", "Hi, You are receiving this email because you requested a password reset.  Use the the password reset code below to reset your password. \n \n Password reset code: " + passwordResetCode)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPassword)

	if err := dialer.DialAndSend(message); err != nil {
		log.Fatal(err)
	}
	
}