package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GenericErrorHandler(c *fiber.Ctx, err error) error {
	fmt.Println(err)
	switch err {
	case gorm.ErrRecordNotFound:
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Record not found",
		})
	case fiber.ErrBadRequest:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
}
