package models

import (
	"log"

	"gorm.io/gorm"
)

// Migrate выполняет миграции для всех моделей
func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&YoutubeChannel{},
		&YoutubeVideo{},
		&MemeTemplate{},
		&RedditPost{},
		&SubReddits{},
		&MemeTopic{},
		&GeneratedMeme{},
		&OurRedditPost{},
		&User{},
		&OurInstagramPost{},
	)
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}
}
