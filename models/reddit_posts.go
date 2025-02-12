package models

import (
	"time"

	"gorm.io/datatypes"
)

type RedditPost struct {
	ID        uint           `gorm:"primaryKey"`
	PostLink  string         `gorm:"column:post_link;type:text;not null"`
	PostTitle string         `gorm:"column:post_title;type:text;not null"`
	PostText  string         `gorm:"column:post_text;type:text"`
	Upvotes   int            `gorm:"column:upvotes"`
	Comments  int            `gorm:"column:comments"`
	PicUrls   datatypes.JSON `gorm:"column:pic_urls;type:jsonb"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	// Внешний ключ для связи с сабреддитом
	SubRedditID uint        `gorm:"column:post_subreddit"`
	SubReddit   *SubReddits `gorm:"foreignKey:SubRedditID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Один Reddit-пост может быть привязан к нескольким темам мема
	MemeTopics []MemeTopic `gorm:"foreignKey:RedditPostID"`
}
