package models

import "time"

type OurRedditPost struct {
	ID       uint   `gorm:"primaryKey"`
	PostLink string `gorm:"column:post_link;type:text;not null"`

	// Ссылка на сгенерированный мем (generated_memes)
	MemeID uint          `gorm:"column:meme_id"`
	Meme   GeneratedMeme `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Karma    int `gorm:"column:karma"`
	Comments int `gorm:"column:comments"`

	// Ссылка на пользователя (users)
	CreatorID uint `gorm:"column:creator"`
	Creator   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
