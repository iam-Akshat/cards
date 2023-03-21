package models_tests

import (
	"testing"

	models "github.com/iam-Akshat/cards/models"
)

func TestCreateNewDeck(t *testing.T) {
	var newDeckTests = []struct {
		shuffled    bool
		cardsString []string
		cardsLength int
		cards       []models.Card
	}{
		{false, []string{}, 52, []models.Card{
			{Suit: models.Spades, Value: models.King, Code: "KS"},
			{Suit: models.Spades, Value: models.Queen, Code: "QS"},
			{Suit: models.Spades, Value: models.Jack, Code: "JS"},
			{Suit: models.Spades, Value: models.Ten, Code: "10S"},
			{Suit: models.Spades, Value: models.Nine, Code: "9S"},
			{Suit: models.Spades, Value: models.Eight, Code: "8S"},
			{Suit: models.Spades, Value: models.Seven, Code: "7S"},
			{Suit: models.Spades, Value: models.Six, Code: "6S"},
			{Suit: models.Spades, Value: models.Five, Code: "5S"},
			{Suit: models.Spades, Value: models.Four, Code: "4S"},
			{Suit: models.Spades, Value: models.Three, Code: "3S"},
			{Suit: models.Spades, Value: models.Two, Code: "2S"},
			{Suit: models.Spades, Value: models.Ace, Code: "AS"},
			{Suit: models.Clubs, Value: models.King, Code: "KC"},
			{Suit: models.Clubs, Value: models.Queen, Code: "QC"},
			{Suit: models.Clubs, Value: models.Jack, Code: "JC"},
			{Suit: models.Clubs, Value: models.Ten, Code: "10C"},
			{Suit: models.Clubs, Value: models.Nine, Code: "9C"},
			{Suit: models.Clubs, Value: models.Eight, Code: "8C"},
			{Suit: models.Clubs, Value: models.Seven, Code: "7C"},
			{Suit: models.Clubs, Value: models.Six, Code: "6C"},
			{Suit: models.Clubs, Value: models.Five, Code: "5C"},
			{Suit: models.Clubs, Value: models.Four, Code: "4C"},
			{Suit: models.Clubs, Value: models.Three, Code: "3C"},
			{Suit: models.Clubs, Value: models.Two, Code: "2C"},
			{Suit: models.Clubs, Value: models.Ace, Code: "AC"},
			{Suit: models.Hearts, Value: models.King, Code: "KH"},
			{Suit: models.Hearts, Value: models.Queen, Code: "QH"},
			{Suit: models.Hearts, Value: models.Jack, Code: "JH"},
			{Suit: models.Hearts, Value: models.Ten, Code: "10H"},
			{Suit: models.Hearts, Value: models.Nine, Code: "9H"},
			{Suit: models.Hearts, Value: models.Eight, Code: "8H"},
			{Suit: models.Hearts, Value: models.Seven, Code: "7H"},
			{Suit: models.Hearts, Value: models.Six, Code: "6H"},
			{Suit: models.Hearts, Value: models.Five, Code: "5H"},
			{Suit: models.Hearts, Value: models.Four, Code: "4H"},
			{Suit: models.Hearts, Value: models.Three, Code: "3H"},
			{Suit: models.Hearts, Value: models.Two, Code: "2H"},
			{Suit: models.Hearts, Value: models.Ace, Code: "AH"},
			{Suit: models.Diamonds, Value: models.King, Code: "KD"},
			{Suit: models.Diamonds, Value: models.Queen, Code: "QD"},
			{Suit: models.Diamonds, Value: models.Jack, Code: "JD"},
			{Suit: models.Diamonds, Value: models.Ten, Code: "10D"},
			{Suit: models.Diamonds, Value: models.Nine, Code: "9D"},
			{Suit: models.Diamonds, Value: models.Eight, Code: "8D"},
			{Suit: models.Diamonds, Value: models.Seven, Code: "7D"},
			{Suit: models.Diamonds, Value: models.Six, Code: "6D"},
			{Suit: models.Diamonds, Value: models.Five, Code: "5D"},
			{Suit: models.Diamonds, Value: models.Four, Code: "4D"},
			{Suit: models.Diamonds, Value: models.Three, Code: "3D"},
			{Suit: models.Diamonds, Value: models.Two, Code: "2D"},
			{Suit: models.Diamonds, Value: models.Ace, Code: "AD"},
		}},
		{false, []string{"AS", "AH", "AC", "AD", "10D"}, 5, []models.Card{
			{Suit: models.Spades, Value: models.Ace, Code: "AS"},
			{Suit: models.Hearts, Value: models.Ace, Code: "AH"},
			{Suit: models.Clubs, Value: models.Ace, Code: "AC"},
			{Suit: models.Diamonds, Value: models.Ace, Code: "AD"},
			{Suit: models.Diamonds, Value: models.Ten, Code: "10D"},
		}},
	}

	for _, tt := range newDeckTests {
		deck := models.CreateNewDeck(tt.shuffled, tt.cardsString)
		if len(deck.Cards.Data) != tt.cardsLength {
			t.Errorf("Expected %d cards, but got %d", tt.cardsLength, len(deck.Cards.Data))
		}
		for i, card := range deck.Cards.Data {
			if card != tt.cards[i] {
				t.Errorf("Expected %s, but got %s", tt.cards[i], card)
			}
		}
	}
}

func TestDeckDraw(t *testing.T) {
	deck := models.CreateNewDeck(false, []string{})
	drawedCards := deck.DrawCards(5)
	if len(deck.Cards.Data) != 47 {
		t.Errorf("Expected 47 cards, but got %d", len(deck.Cards.Data))
	}
	expectedCards := []models.Card{
		{Suit: models.Spades, Value: models.King, Code: "KS"},
		{Suit: models.Spades, Value: models.Queen, Code: "QS"},
		{Suit: models.Spades, Value: models.Jack, Code: "JS"},
		{Suit: models.Spades, Value: models.Ten, Code: "10S"},
		{Suit: models.Spades, Value: models.Nine, Code: "9S"},
	}
	for i, card := range drawedCards {
		if card != expectedCards[i] {
			t.Errorf("Expected %s, but got %s", expectedCards[i], card)
		}
	}
}
