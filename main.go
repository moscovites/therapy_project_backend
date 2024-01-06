package main

import (
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

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
}

func serverApplication() {
	router := gin.Default()
	publicRoutes := router.Group("/waitlist")
	publicRoutes.POST("/member", controllers.CreateWaitlistMember)
	
	router.Run(":8000")
	fmt.Println("server running on port 8000")
}