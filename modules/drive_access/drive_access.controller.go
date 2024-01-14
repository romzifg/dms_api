package drive_access

import (
	"dms_api/database"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	var drive_access []DriveAccess

	result := database.DB.Find(&drive_access)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message":    "Bad Request",
			"data":       nil,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       drive_access,
	})
}

func Store(c *fiber.Ctx) error {
	var payload []DriveAccess

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})
	}

	result := database.DB.Create(&payload)
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
		"data": payload,
	})
}

func Update(c *fiber.Ctx) error {
	var drive_access DriveAccess
	var dto DriveAccessUpdateDto
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

	if err = database.DB.First(&drive_access, id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"statusCode": http.StatusNotFound,
			"message": "Bad Request",
			"data": nil,
		})
	}

	result := database.DB.Model(&drive_access).Where("drive_id = ?", id).Updates(&dto)
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
		"data": drive_access,
	})
}

func DeleteData(c *fiber.Ctx) error {
	var drive_access *DriveAccess
	id := c.Params("id")

	deletePayload := DriveAccess{
		DeletedAt: time.Now(),
	}

	result := database.DB.Model(&drive_access).Where("drive_id = ?", id).Updates(&deletePayload)
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