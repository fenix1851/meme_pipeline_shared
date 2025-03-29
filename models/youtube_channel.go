package models

import "time"

type YoutubeChannel struct {
	ID           int       `gorm:"column:id;primaryKey;autoIncrement"`
	Link   string    `gorm:"column:link"`
	Subscribers  int       `gorm:"column:subscribers"`
	OriginalData string    `gorm:"column:original_data"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
