package main

import (
	"github.com/gin-contrib/cors"
	"core/database"
	"core/models"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"github.com/gin-gonic/gin"
	"core/controllers"
)

func main () {
	loadEnv()
	loadDatabase()
	serverApplication()
}

func loadDatabase() {
	
	database.Connect()
	database.Database.AutoMigrate(&models.Waitlist{})
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.PatientProfile{})
	database.Database.AutoMigrate(&models.TherapistProfile{})

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
}

func serverApplication() {
    router := gin.Default()

    // CORS middleware
    router.Use(cors.Default())

    
	publicRoutes := router.Group("/waitlist")
	publicRoutes.POST("/member", controllers.CreateWaitlistMember)

	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.GET("verify-email/:code", controllers.VerifyEmail)
	publicRoutes.POST("/login", controllers.Login)
	publicRoutes.POST("/password-reset", controllers.PasswordReset)
	publicRoutes.GET("/:password-reset-code", controllers.PasswordResetConfirm)
	publicRoutes.POST("/:password-reset-code", controllers.CreateNewPassword)

	onBoardingRoutes := router.Group("/onboarding")

	onBoardingRoutes.POST("/patient", controllers.CreatePatientProfile)
	onBoardingRoutes.POST("/therapist", controllers.CreateTherapistProfile)
	onBoardingRoutes.PUT("/patient", controllers.UpdatePatientProfile)
	onBoardingRoutes.PUT("/therapist", controllers.UpdateTherapistProfile)

	
	
	router.Run(":8000")
	fmt.Println("server running on port 8000")
}