package auth

import (
	"dms_api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	api := app.Group("/api/v1/auth")
	api.Use(middlewares.ApiTokenMiddleware)

	api.Post("/login", Login)
}