package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iam-Akshat/cards/handlers"
	"github.com/iam-Akshat/cards/models"
	"github.com/iam-Akshat/cards/services"
	"gorm.io/gorm"
)

// TODO:- Remove Migration from here

func SetupDeckRoutes(app fiber.Router, db *gorm.DB) {
	db.AutoMigrate(&models.Deck{})
	deckService := services.NewDeckService(db)
	deckHandler := handlers.NewDeckHandler(deckService)

	app.Post("/deck", deckHandler.CreateDeck)
	app.Get("/deck/:deck_id", deckHandler.GetDeck)
	app.Post("/deck/:deck_id/draw", deckHandler.DrawCards)
}
