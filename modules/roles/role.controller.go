package roles

import (
	"dms_api/database"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllData(c *fiber.Ctx) error {
	var roles []Role

	database.DB.Find(&roles)
	c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": roles,
	})

	return nil
}

func Store(c *fiber.Ctx) error {
	var role Role

	if err := c.BodyParser(&role); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})
	}

	database.DB.Create(&role)
	c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": role,
	})

	return nil
}