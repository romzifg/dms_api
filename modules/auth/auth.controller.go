package auth

import (
	"dms_api/database"
	"dms_api/helpers"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	var authM Auth
	var loginDto AuthLoginDto

	err := c.BodyParser(&loginDto)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})
	}

	err = database.DB.Where("email = ?", loginDto.Email).Find(&authM).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"statusCode": http.StatusNotFound,
				"message": "Not Found",
				"data": nil,
			})
		default: 
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"statusCode": http.StatusBadRequest,
				"message": "Bad Request",
				"data": nil,
			})
		}
	}

	_, errCompare := helpers.ComparePassword(authM.Password, loginDto.Password)
	if errCompare != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": http.StatusUnauthorized,
			"message": "Unauthorized",
			"data": nil,
		})
	}

	token, err := helpers.GenerateToken(int(authM.UserId))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"statusCode": http.StatusUnauthorized,
			"message": "Unauthorized",
			"data": nil,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": http.StatusOK,
		"message": "Success",
		"token": token,
	})
}