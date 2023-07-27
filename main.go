package main

import (
	"clockwork-server/config"
	"clockwork-server/infra/database"
	"clockwork-server/interfaces/api/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	db := database.NewDatabase()
	r := router.NewRouter(db.Connect())

	port := config.GetConfig().Server.Port
	r.RegisterAPI().Run(":" + port)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
}
