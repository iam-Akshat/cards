package models

import (
	"math/rand"
	"time"
)

type CardSuit string

const (
	Spades   CardSuit = "SPADES"
	Clubs    CardSuit = "CLUBS"
	Hearts   CardSuit = "HEARTS"
	Diamonds CardSuit = "DIAMONDS"
)

type CardValue string

const (
	King  CardValue = "KING"
	Queen CardValue = "QUEEN"
	Jack  CardValue = "JACK"
	Ten   CardValue = "10"
	Nine  CardValue = "9"
	Eight CardValue = "8"
	Seven CardValue = "7"
	Six   CardValue = "6"
	Five  CardValue = "5"
	Four  CardValue = "4"
	Three CardValue = "3"
	Two   CardValue = "2"
	Ace   CardValue = "ACE"
)

type CardCode string

type Card struct {
	Suit  CardSuit  `json:"suit"`
	Value CardValue `json:"value"`
	Code  CardCode  `json:"code"`
}

// should change to static
var SequencedAllCards []Card = make([]Card, 0, 52)
var CodeToCard map[CardCode]Card = make(map[CardCode]Card)

func GetCardCode(suit CardSuit, value CardValue) CardCode {
	specialCards := map[CardValue]string{
		King:  "K",
		Queen: "Q",
		Jack:  "J",
		Ace:   "A",
		"10":  "10",
	}
	if specialCard, ok := specialCards[value]; ok {
		return CardCode(specialCard) + CardCode(suit[0])
	}
	return CardCode(value[0]) + CardCode(suit[0])
}

func NewCard(suit CardSuit, value CardValue) Card {
	return Card{
		Suit:  suit,
		Value: value,
		Code:  GetCardCode(suit, value),
	}
}

func GetCardFromCode(code CardCode) Card {
	return CodeToCard[code]
}

func ShuffleCards(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards
}

func init() {
	for _, suit := range []CardSuit{Spades, Clubs, Hearts, Diamonds} {
		for _, value := range []CardValue{King, Queen, Jack, Ten, Nine, Eight, Seven, Six, Five, Four, Three, Two, Ace} {
			card := NewCard(suit, value)
			CodeToCard[card.Code] = card
			SequencedAllCards = append(SequencedAllCards, card)
		}
	}
}
