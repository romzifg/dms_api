package role_permissions

import (
	"dms_api/database"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	var permission []RolePermission
	query := c.Query("role_id") 

	result := database.DB.Where("role_id = ?", query).Find(&permission)
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
		"data":       permission,
	})
}

func Store(c *fiber.Ctx) error {
	var inputDto *RolePermissionInputDto
	var payload []RolePermission

	if err := c.BodyParser(&inputDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})
	}

	permissionsData := inputDto.Permission

	for _, input := range permissionsData {
		payload = append(payload, RolePermission{
			RoleId: int(inputDto.RoleId),
			Code: input.Code,
			Description: input.Description,
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
	var inputDto *RolePermissionInputDto
	var payload []RolePermission
	id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Model(&payload).Where("role_id = ?", id).Delete(&payload)
	if err := c.BodyParser(&inputDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": "Bad Request",
			"data": nil,
		})
	}

	permissionsData := inputDto.Permission

	for _, input := range permissionsData {
		payload = append(payload, RolePermission{
			RoleId: int(id),
			Code: input.Code,
			Description: input.Description,
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