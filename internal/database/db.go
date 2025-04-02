package database

import (
	"fmt"
	"github.com/allang-4779/financer/internal/configuration"
	"github.com/allang-4779/financer/internal/models"
	customqueries "github.com/allang-4779/financer/internal/models/custom-queries"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB(cfg *configuration.Configuration) {
	dsn := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPort, cfg.DBPassword, cfg.DBName, cfg.SSLMode)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Could not initialize database connection")
	}
	log.Println("Database connected successfully")
}

func MigrateDB() {
	err := DB.AutoMigrate(models.SystemUser{},
		models.Account{}, models.Category{}, models.Transaction{}, models.Budget{})
	if err != nil {
		log.Fatal("Migration failed:", err)

	}
	customqueries.CreateTriggerIfNotExists(DB)

	log.Println("Database migrated successfully")
}
