package upload

import (
	"dms_api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	api := app.Group("/api/v1/upload")
	api.Use(middlewares.ApiTokenMiddleware)

	api.Post("/file", UploadFile)
	api.Post("/files", MultipleUpload)
}