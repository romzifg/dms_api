package roles

import (
	"dms_api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	api := app.Group("/api/v1/role")
	api.Use(middlewares.ApiTokenMiddleware)

	api.Get("/", GetAll)
	api.Post("/", Store)
	api.Put("/:id", Update)
	api.Delete("/:id", DeleteData)
}