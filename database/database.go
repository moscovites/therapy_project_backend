package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error
	// host := os.Getenv("DB_HOST")
	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// databaseName := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("postgresql://postgres:dWGkIkpFvlwZigvBZBJcHZQWXWJgitNe@monorail.proxy.rlwy.net:50876/railway")
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}
