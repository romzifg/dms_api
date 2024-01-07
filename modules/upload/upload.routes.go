package upload

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	api := app.Group("/api/v1/upload")

	api.Post("/", UploadFile)
}