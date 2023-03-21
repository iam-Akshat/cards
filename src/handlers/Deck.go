package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/iam-Akshat/cards/services"
	"github.com/iam-Akshat/cards/utils"
)

type DeckHandler struct {
	DeckService services.DeckService
}

func NewDeckHandler(deckService *services.DeckService) *DeckHandler {
	return &DeckHandler{
		DeckService: *deckService,
	}
}

type CreateDeckRequest struct {
	Shuffled bool     `json:"shuffled"`
	Cards    []string `json:"cards"`
}

type DrawCardsRequest struct {
	Count int `json:"count"`
}

// TODO: Use a validation library for better validation and error handling
func (handler *DeckHandler) CreateDeck(ctx *fiber.Ctx) error {
	reqBody := new(CreateDeckRequest)
	if err := ctx.BodyParser(reqBody); err != nil {
		return utils.GenericErrorHandler(ctx, fiber.ErrBadRequest)
	}

	deck, err := handler.DeckService.CreateDeck(reqBody.Shuffled, reqBody.Cards)
	if err != nil {
		return utils.GenericErrorHandler(ctx, err)
	}

	return ctx.Status(http.StatusCreated).JSON(deck.ToCreateResponse())

}

func (handler *DeckHandler) GetDeck(ctx *fiber.Ctx) error {
	deckId := ctx.Params("deck_id")

	deck, err := handler.DeckService.GetDeck(uuid.MustParse(deckId))
	if err != nil {
		return utils.GenericErrorHandler(ctx, err)
	}

	return ctx.JSON(deck.ToGetResponse())
}

func (handler *DeckHandler) DrawCards(ctx *fiber.Ctx) error {
	deckId := ctx.Params("deck_id")

	reqBody := new(DrawCardsRequest)
	if err := ctx.BodyParser(reqBody); err != nil {
		return utils.GenericErrorHandler(ctx, fiber.ErrBadRequest)
	}

	count := reqBody.Count
	if count <= 0 || count > 52 || string(rune(count)) == "" {
		return utils.GenericErrorHandler(ctx, fiber.ErrBadRequest)
	}

	deck, err := handler.DeckService.DrawCards(uuid.MustParse(deckId), count)
	if err != nil {
		return utils.GenericErrorHandler(ctx, err)
	}

	return ctx.JSON(deck)
}
