package db

import (
	"final_project_cloud_2024/config"
	"final_project_cloud_2024/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectAndMigrate() *gorm.DB {
	config.LoadEnv()

	host := config.GetEnv("DB_HOST", "localhost")
	user := config.GetEnv("DB_USER", "postgres")
	password := config.GetEnv("DB_PASSWORD", "postgres")
	dbname := config.GetEnv("DB_NAME", "eventdb")
	port := config.GetEnv("DB_PORT", "5432")
	sslmode := config.GetEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Connected to the database successfully.")

	// Run migrations
	err = db.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Registration{},
		&models.Notification{},
		&models.Payment{},
		&models.Venue{},
		&models.Ticket{},
		&models.Review{},
		&models.Category{},
		&models.EventCategory{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Database migrated successfully.")

	return db
}
