package roles

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	api := app.Group("/api/v1/role")

	api.Get("/", GetAllData)
	api.Post("/", Store)
	api.Put("/:id", Update)
	api.Delete("/:id", DeleteData)
}