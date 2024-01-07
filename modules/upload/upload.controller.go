package upload

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")

	if err != nil {
		return err
	}

	destination := fmt.Sprintf("./public/files/%s", file.Filename)
	if err = c.SaveFile(file, destination); err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"data": destination,
	})
}