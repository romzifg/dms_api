package roles

import (
	"dms_api/database"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	var roles []Role

	database.DB.Find(&roles)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": roles,
	})
}

func Store(c *fiber.Ctx) error {
	var role *Role

	if err := c.BodyParser(&role); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})
	}

	newRole := Role{
		RoleName: role.RoleName,
		Status: role.Status,
	}

	result := database.DB.Create(&newRole)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": role,
	})
}

func Update(c *fiber.Ctx) error {
	var role Role
	var dto UpdateRoleDto
	id := c.Params("id")
	_, err := strconv.Atoi(id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})
	}

	if err = c.BodyParser(&dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})
	}

	if err = database.DB.First(&role, id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"statusCode": http.StatusNotFound,
			"message": "Bad Request",
			"data": nil,
		})
	}

	result := database.DB.Model(&role).Where("role_id = ?", id).Updates(&dto)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": role,
	})
}

func DeleteData(c *fiber.Ctx) error {
	var role *Role
	id := c.Params("id")

	deletePayload := Role{
		DeletedAt: time.Now(),
	}

	result := database.DB.Model(&role).Where("role_id = ?", id).Updates(&deletePayload)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": id,
	})
}