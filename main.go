package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	app := App{}

	// Load env file.
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("rror loading .env file")
	}

	app.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	app.Run(":8080")
}
