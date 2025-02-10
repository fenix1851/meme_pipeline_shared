package models

import (
	"log"

	"gorm.io/gorm"
)

// Migrate выполняет миграции для всех моделей
func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&RedditPost{},
		&MemeTemplate{},
		&User{},
		&Meme{},
		&MemeKarma{},
		&MemeComment{},
	)
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}
}
