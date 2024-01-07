package upload

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": err.Error(),
			"data": nil,
		})
	}

	destination := fmt.Sprintf("./public/files/%s", file.Filename)
	if err = c.SaveFile(file, destination); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": err.Error(),
			"data": nil,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": destination,
	})
}

func MultipleUpload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	var listDest []string;

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"statusCode": http.StatusBadRequest,
			"message": err.Error(),
			"data": nil,
		})
	}

	for _, fileHeaders := range form.File {
		for _, fileHeader := range fileHeaders {
			destination := fmt.Sprintf("./public/files/%s", fileHeader.Filename)
			if err = c.SaveFile(fileHeader, destination); err != nil {
					return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"statusCode": http.StatusBadRequest,
					"message": err.Error(),
					"data": nil,
				})
			}

			listDest = append(listDest, destination)
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": http.StatusOK,
		"message": "Success",
		"data": listDest,
	})
}