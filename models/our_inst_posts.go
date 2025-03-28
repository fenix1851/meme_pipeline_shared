package models

import "time"

type OurInstagramPost struct {
	ID          uint   `gorm:"primaryKey"`
	PostLink    string `gorm:"column:post_link;type:text;not null"`
	ImageURL    string `gorm:"column:image_url;type:text;not null"`
	Description string `gorm:"column:description;type:text"`

	// Ссылка на сгенерированный мем (generated_memes)
	MemeID uint          `gorm:"column:meme_id"`
	Meme   GeneratedMeme `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Likes    int `gorm:"column:likes"`
	Comments int `gorm:"column:comments"`

	// TODO: Ссылка на пользователя (users)
	CreatorID uint `gorm:"column:creator"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
