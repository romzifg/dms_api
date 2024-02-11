package middlewares

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func ApiTokenMiddleware(c *fiber.Ctx) error {
	requiredToken := os.Getenv("REQUIRED_TOKEN")
	token := c.Get("api_token")

	if token == "" || token != requiredToken {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": http.StatusUnauthorized,
			"message": "Unauthorized",
		})
	}

	return c.Next()
}