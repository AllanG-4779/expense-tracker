package main

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"

	"github.com/allang-4779/financer/internal/configuration"
	"github.com/allang-4779/financer/internal/database"

	"github.com/allang-4779/financer/internal/routes"
	"github.com/allang-4779/financer/internal/util"
	"github.com/gin-gonic/gin"
)

func main() {
	start := time.Now()

	cfg := configuration.InitializeConfiguration()
	database.InitDB(cfg)
	// LOAD RSA KEYS
	util.LoadKeys("private.pem", "public.pem")

	end := time.Now()

	log.Printf("Intialization took %v seconds", (end.UnixMilli()-start.UnixMilli())/1000)
	database.MigrateDB()

	// Setup GIN server

	r := gin.Default()

	// REGISTER ROUTES
	routes.UserRoutesSetup(r)
	defer func(DB *sqlx.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal("Failed to close database connection")
		}
	}(database.DB)

	err := r.Run(":9097")
	if err != nil {
		log.Fatal("Failed to start 	HTTP SERVER")
	}
	log.Println("Server started successfully")

}
