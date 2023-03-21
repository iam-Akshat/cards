package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	routes "github.com/iam-Akshat/cards/routes"

	db "github.com/iam-Akshat/cards/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	dbConfig := db.GetDBConfigFromEnv()
	db, err := db.NewDatabaseConnection(dbConfig)
	if err != nil {
		panic(err)
	}
	app.Use(logger.New())
	routes.SetupRoutes(app, db)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ok fffthids works")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal((app.Listen(":" + port)))
}
