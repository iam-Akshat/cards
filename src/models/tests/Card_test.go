package models_tests

import (
	"testing"

	models "github.com/iam-Akshat/cards/models"
)

var testCases = []struct {
	suit  models.CardSuit
	value models.CardValue
	code  models.CardCode
}{
	{models.Spades, models.Ace, "AS"},
	{models.Hearts, models.King, "KH"},
	{models.Clubs, models.Queen, "QC"},
	{models.Diamonds, models.Jack, "JD"},
	{models.Spades, models.Ten, "10S"},
}

func TestNewCard(t *testing.T) {

	for _, testCase := range testCases {
		card := models.NewCard(testCase.suit, testCase.value)
		if card.Suit != testCase.suit {
			t.Errorf("Expected suit to be %s, got %s", testCase.suit, card.Suit)
		}
		if card.Value != testCase.value {
			t.Errorf("Expected value to be %s, got %s", testCase.value, card.Value)
		}
		if card.Code != testCase.code {
			t.Errorf("Expected code to be %s, got %s", testCase.code, card.Code)
		}
	}

}

func TestGetCardFromCode(t *testing.T) {

	for _, testCase := range testCases {
		card := models.GetCardFromCode(testCase.code)
		if card.Suit != testCase.suit {
			t.Errorf("Expected suit to be %s, got %s", testCase.suit, card.Suit)
		}
		if card.Value != testCase.value {
			t.Errorf("Expected value to be %s, got %s", testCase.value, card.Value)
		}
		if card.Code != testCase.code {
			t.Errorf("Expected code to be %s, got %s", testCase.code, card.Code)
		}
	}

}
