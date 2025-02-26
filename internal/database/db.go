package database

import (
	"fmt"
	"log"

	"github.com/allang-4779/financer/internal/configuration"
	"github.com/allang-4779/financer/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *configuration.Configuration){
	dsn := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=disable",
	cfg.DBHost, cfg.DBUser, cfg.DBPort,cfg.DBPassword,cfg.DBName) 

	var error error

	DB, error  = gorm.Open(postgres.Open(dsn),&gorm.Config{} )

	if  error != nil {
		log.Fatal("Could not initialize database connection")
	}
    log.Println("Database connected successfully")
}


func MigrateDB() {
    err := DB.AutoMigrate(models.SystemUser{}, &models.LoginAccount{})
    if err != nil {
        log.Fatal("Migration failed:", err)
    }
    log.Println("Database migrated successfully")
}