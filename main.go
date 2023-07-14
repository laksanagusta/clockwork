package main

import (
	"clockwork-server/app"
	"clockwork-server/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	db := database.NewDatabase()
	r := app.NewRouter(db.Connect())

	port := os.Getenv("PORT")
	r.RegisterAPI().Run(":" + port)
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
