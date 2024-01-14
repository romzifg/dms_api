package drive

import (
	"dms_api/database"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	var drive []Drive

	result := database.DB.Find(&drive)
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
		"data":       drive,
	})
}

func Store(c *fiber.Ctx) error {
	var payload Drive

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
	var drive Drive
	var dto DriveUpdateDto
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

	if err = database.DB.First(&drive, id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"statusCode": http.StatusNotFound,
			"message": "Bad Request",
			"data": nil,
		})
	}

	result := database.DB.Model(&drive).Where("drive_id = ?", id).Updates(&dto)
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
		"data": drive,
	})
}

func DeleteData(c *fiber.Ctx) error {
	var drive *Drive
	id := c.Params("id")

	deletePayload := Drive{
		DeletedAt: time.Now(),
	}

	result := database.DB.Model(&drive).Where("drive_id = ?", id).Updates(&deletePayload)
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