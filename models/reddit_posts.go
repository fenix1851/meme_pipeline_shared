package models

import "time"

type RedditPost struct {
	ID        uint   `gorm:"primaryKey"`
	PostLink  string `gorm:"column:post_link;type:text;not null"`
	PostTitle string `gorm:"column:post_title;type:text;not null"`
	PostText  string `gorm:"column:post_text;type:text"`
	Karma     int    `gorm:"column:karma"`
	// Связь с reddit_treads
	RedditTreadID uint        `gorm:"column:post_tread"`
	RedditTread   RedditTread `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments      int         `gorm:"column:comments"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// Один Reddit-пост может быть привязан к нескольким темам мема
	MemeTopics []MemeTopic `gorm:"foreignKey:RedditPostID"`
}
