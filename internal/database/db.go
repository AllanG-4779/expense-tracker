package database

import (
	"database/sql"
	"fmt"
	"github.com/allang-4779/financer/internal/configuration"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

var DB *sql.DB

func InitDB(cfg *configuration.Configuration) {
	dsn := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPort, cfg.DBPassword, cfg.DBName)

	var err error

	DB, err = sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("Could not initialize database connection")
	}
	log.Println("Database connected successfully")
}

func MigrateDB() {
	driver, err := postgres.WithInstance(DB, &postgres.Config{})

	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal("Could not load migrations: ", err)
	}
	if err := m.Up(); err != nil {
		log.Fatal("Migration failed: ", err)
	}

	log.Println("Database migrated successfully")
}
