package config

import (
	"fmt"
	"log"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"time"
	// "ihsansolusi-account/apps/databases/models"
)

var DB *gorm.DB

func InitDB() {
	if err := godotenv.Load(); err != nil {
        log.Println("Warning: .env file not found, using system environment variables")
    }

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	/*
	//* Auto migrate model
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}
	*/

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Database connected successfully.")
}

func GetDB() *gorm.DB {
	return DB
}