package models

import (
	"time"

	"gorm.io/datatypes"
)

type GeneratedMeme struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	// Ссылка на шаблон мема (meme_templates)
	MemeTemplateID uint         `gorm:"column:meme_pic"` // foreign key: meme_templates.id
	MemeTemplate   MemeTemplate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// Ссылка на тему мема (meme_topics)
	MemeTopicID uint      `gorm:"column:meme_topic"` // foreign key: meme_topics.id
	MemeTopic   MemeTopic `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TextContent datatypes.JSON `gorm:"column:text_content;type:jsonb"`

	// Один сгенерированный мем может быть привязан к нескольким нашим Reddit-постам
	OurRedditPosts []OurRedditPost `gorm:"foreignKey:MemeID"`
}
