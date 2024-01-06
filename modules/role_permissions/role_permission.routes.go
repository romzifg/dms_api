package role_permissions

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	api := app.Group("/api/v1/role-permission")

	api.Get("/", GetAll)
	api.Post("/", Store)
}