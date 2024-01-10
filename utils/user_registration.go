package utils

import (
	"fmt"
	"core/models"
	"core/database"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
	"log"
	"gorm.io/gorm"
)

func IsEmailUnique (email string) bool {
	var existingUsers models.User
	result := database.Database.Where("email = ?", email).First(&existingUsers)
	return 	result.Error == gorm.ErrRecordNotFound
}

func GenerateVerificationCode () string {
	code := uuid.New().String()[:8]
	return code
}

func FindUserByVerificationCode (code string) (*models.User, error) {
	var user models.User
	result := database.Database.Where("verification_code = ?", code).First(&user)
	
	return &user, result.Error
}

func SendVerificationEmail (email, verificationCode string) {
	senderEmail := "stakshare.com@gmail.com"
	senderPassword := "ufbzqxrvvuwrjbuh"

	recipientEmail := email

	message := gomail.NewMessage()
	message.SetHeader("From", senderEmail)
	message.SetHeader("To", recipientEmail)
	message.SetHeader("Subject", "Email Verification, Stakshare.")
	message.SetBody("text/plain", "Hi, You are receiving this email because we want to verify if this email address is indeed yours.Kindly use the 8 digit code below to verify your email address. \n \n Verification code: " + verificationCode)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPassword)

	if err := dialer.DialAndSend(message); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email sent successfully")
}
