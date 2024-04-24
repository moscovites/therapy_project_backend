package main

import (
	"core/controllers"
	"core/database"
	"core/middleware"
	"core/models"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
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
	database.Database.AutoMigrate(&models.Appointment{})

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

	router.POST("/login", controllers.Login)

	protectedRoutes := router.Group("/user")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	protectedRoutes.POST("/apointment", controllers.CreateAppointment)
	protectedRoutes.PUT("/apointment", controllers.UpdateAppointment)
	protectedRoutes.GET("/apointment", controllers.GetAllAppointments)

	// router.GET("/ws", controllers.HandleConnections)

	// // Start a goroutine to handle incoming messages and broadcast them to clients
	// go controllers.HandleMessages()

	publicRoutes := router.Group("/waitlist")
	publicRoutes.POST("/member", controllers.CreateWaitlistMember)

	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.GET("verify-email/:code", controllers.VerifyEmail)

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
