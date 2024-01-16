package drive_access

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	api := app.Group("/api/v1/drive-access")

	api.Get("/", GetAll)
	api.Post("/", Store)
	api.Put("/:id", Update)
	api.Delete("/:id", DeleteData)
}