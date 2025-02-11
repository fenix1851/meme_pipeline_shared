package models

import (
	"log"

	"gorm.io/gorm"
)

// Migrate выполняет миграции для всех моделей
func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&MemeTemplate{},
		&RedditPost{},
		&RedditTread{},
		&MemeTopic{},
		&GeneratedMeme{},
		&OurRedditPost{},
		&User{},
	)
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}
}
