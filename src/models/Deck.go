package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// cards is of type `jsonb` in postgres
type Deck struct {
	gorm.Model
	Id        uuid.UUID                  `json:"deck_id" gorm:"type:uuid;primaryKey"`
	Shuffled  bool                       `json:"shuffled"`
	Remaining int                        `json:"remaining"`
	Cards     datatypes.JSONType[[]Card] `json:"cards"`
}

// these ideally should be in a separate file
type CreateDeckResponse struct {
	DeckId    uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
}

type GetDeckResponse struct {
	DeckId    uuid.UUID                  `json:"deck_id"`
	Shuffled  bool                       `json:"shuffled"`
	Remaining int                        `json:"remaining"`
	Cards     datatypes.JSONType[[]Card] `json:"cards"`
}

type DrawFromDeckResponse struct {
	Cards []Card `json:"cards"`
}

func CreateNewDeck(shuffled bool, cards []string) Deck {
	var cardsObj []Card = []Card{}
	for _, card := range cards {
		cardsObj = append(cardsObj, GetCardFromCode(CardCode(card)))
	}

	// case when no cards are passed, set the default deck
	if len(cards) == 0 {
		cardsObj = append(cardsObj, SequencedAllCards...)
	}
	if shuffled {
		cardsObj = ShuffleCards(cardsObj)
	}

	return Deck{
		Id:        uuid.New(),
		Shuffled:  shuffled,
		Remaining: len(cardsObj),
		Cards:     datatypes.JSONType[[]Card]{Data: cardsObj},
	}
}

func (d *Deck) DrawCards(count int) []Card {
	cards := d.Cards.Data[:count]
	d.Cards.Data = d.Cards.Data[count:]
	d.Remaining = len(d.Cards.Data)
	return cards
}

// move to dtos
func (d *Deck) ToCreateResponse() CreateDeckResponse {
	return CreateDeckResponse{
		DeckId:    d.Id,
		Shuffled:  d.Shuffled,
		Remaining: d.Remaining,
	}
}

func (d *Deck) ToGetResponse() GetDeckResponse {
	return GetDeckResponse{
		DeckId:    d.Id,
		Shuffled:  d.Shuffled,
		Remaining: d.Remaining,
		Cards:     d.Cards,
	}
}

func (d *Deck) ToDrawResponse() DrawFromDeckResponse {
	return DrawFromDeckResponse{
		Cards: d.Cards.Data,
	}
}
