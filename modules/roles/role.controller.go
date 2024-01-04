package roles

import (
	"dms_api/database"
	"net/http"
	"strconv"

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

func Update(c *fiber.Ctx) error {
	var role Role
	var dto UpdateRoleDto
	id := c.Params("id")
	_, err := strconv.Atoi(id)

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})

		return err
	}

	if err = c.BodyParser(&dto); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})

		return err
	}

	if err = database.DB.First(&role, id).Error; err != nil {
		c.Status(http.StatusNotFound).JSON(fiber.Map{
			"statusCode": http.StatusNotFound,
			"message": "Bad Request",
			"data": nil,
		})

		return err
	}

	database.DB.Model(&role).Where("role_id = ?", id).Updates(&dto)
	c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": role,
	})

	return nil
}