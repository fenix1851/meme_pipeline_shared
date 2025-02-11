package models

import "time"

type MemeTopic struct {
	ID uint `gorm:"primaryKey"`

	// Связь с reddit_posts
	RedditPostID uint       `gorm:"column:reddit_post_id"`
	RedditPost   RedditPost `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// Связь с reddit_treads
	RedditTreadID uint        `gorm:"column:post_tread"`
	RedditTread   RedditTread `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Title         string    `gorm:"column:title;type:text;not null"`
	Karma         int       `gorm:"column:karma"`
	FormattedText string    `gorm:"column:formatted_text;type:text"`
	Comments      int       `gorm:"column:comments"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
	// Один топик может иметь несколько сгенерированных мемов
	GeneratedMemes []GeneratedMeme `gorm:"foreignKey:MemeTopicID"`
}
