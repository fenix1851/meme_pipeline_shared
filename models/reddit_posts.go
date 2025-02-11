package models

import "time"

type RedditPost struct {
	ID        uint   `gorm:"primaryKey"`
	PostLink  string `gorm:"column:post_link;type:text;not null"`
	PostTitle string `gorm:"column:post_title;type:text;not null"`
	PostText  string `gorm:"column:post_text;type:text"`
	Upvotes   int    `gorm:"column:upvotes"`
	// Связь с reddit_treads
	SubRedditId uint       `gorm:"column:post_subreddit"`
	SubReddits  SubReddits `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments    int        `gorm:"column:comments"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// Один Reddit-пост может быть привязан к нескольким темам мема
	MemeTopics []MemeTopic `gorm:"foreignKey:RedditPostID"`
}
