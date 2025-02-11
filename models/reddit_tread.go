package models

import "time"

type RedditTread struct {
	ID        uint      `gorm:"primaryKey"`
	TreadName string    `gorm:"column:tread_name;type:text;not null"`
	TreadLink string    `gorm:"column:tread_link;type:text;not null"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// Один тред может быть привязан к нескольким темам мема
	MemeTopics []MemeTopic `gorm:"foreignKey:RedditTreadID"`
}
