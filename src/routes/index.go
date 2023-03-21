package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	group := app.Group("/api/v1")

	SetupDeckRoutes(group, db)
}
