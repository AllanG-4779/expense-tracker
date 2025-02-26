package main

import (
	"log"
	"time"

	"github.com/allang-4779/financer/internal/configuration"
	"github.com/allang-4779/financer/internal/database"
	"github.com/allang-4779/financer/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {	
	start := time.Now();

	cfg := configuration.InitializeConfiguration();
	database.InitDB(cfg)	

	end := time.Now()

	log.Printf("Intialization took %v seconds", (end.UnixMilli() - start.UnixMilli())/1000 )
	database.MigrateDB()

	// Setup GIN server

	r := gin.Default();

	// REGISTER ROUTES
	routes.UserRoutesSetup(r)

	err := r.Run(":9095")
	if err != nil {
		log.Fatal("Failed to start 	HTTP SERVER")
	}
	log.Println("Server started successfully")

	
}

