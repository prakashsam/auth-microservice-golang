package main

import (
	"authservice/config"
	"authservice/db"
	"authservice/route"
	"authservice/utils"
	"log"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	godotenv.Load()
	cfg := config.Load()
	db.InitDBConnection()

	app := iris.New()

	jwtSecret, err := utils.GetSecret("JWT_SECRET", cfg.ProjectID)
	if err != nil {
		log.Fatalf("Failed to get JWT secret: %v", err)
	}

	route.SetupRoutes(app, jwtSecret)
	app.Listen(":8081")
}
