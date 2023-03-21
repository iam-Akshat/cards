package services

import (
	"github.com/google/uuid"
	"github.com/iam-Akshat/cards/models"
	"gorm.io/gorm"
)

type DeckService struct {
	db *gorm.DB
}

func NewDeckService(db *gorm.DB) *DeckService {
	return &DeckService{
		db: db,
	}
}

// Add validation cases like:
// Only 4 valid suits are allowed
// Only 13 valid ranks are allowed
// No duplicate cards are allowed etc.
func (s *DeckService) CreateDeck(shuffled bool, cards []string) (models.Deck, error) {

	var deck models.Deck = models.CreateNewDeck(shuffled, cards)

	result := s.db.Create(&deck)
	if result.Error != nil {
		return models.Deck{}, result.Error
	}
	return deck, nil
}

func (s *DeckService) GetDeck(id uuid.UUID) (models.Deck, error) {

	var deck models.Deck

	result := s.db.First(&deck, id)
	if result.Error != nil {
		return models.Deck{}, result.Error
	}

	return deck, nil
}

func (s *DeckService) DrawCards(id uuid.UUID, count int) ([]models.Card, error) {

	var deck models.Deck

	result := s.db.First(&deck, id)
	if result.Error != nil {
		return []models.Card{}, result.Error
	}

	cards := deck.DrawCards(count)

	result = s.db.Save(&deck)
	if result.Error != nil {
		return []models.Card{}, result.Error
	}

	return cards, nil
}
