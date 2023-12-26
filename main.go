package main

import (
	"dms_api/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName: os.Getenv("GO_APP_NAME"),
	})

	// connection database
	database.DbConnect()

	// static file
	app.Static("/", "./public") 
	app.Listen(":" + os.Getenv("GO_API_PORT"))
}