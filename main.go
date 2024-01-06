package main

import (
	"dms_api/database"
	"dms_api/modules/role_permissions"
	"dms_api/modules/roles"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	// connection database
	database.DbConnect()

	// static file
	app.Static("/", "./public") 

	// Routes
	roles.Routes(app)
	role_permissions.Routes(app)

	log.Fatal(app.Listen(":" + os.Getenv("GO_API_PORT")))
}