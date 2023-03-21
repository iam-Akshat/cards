package testutils

import (
	"github.com/iam-Akshat/cards/models"
	"gorm.io/gorm"
)

func DropAllTables(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Deck{}, &models.Card{})
}
