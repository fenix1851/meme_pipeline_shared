package models

import (
	"time"

	"github.com/lib/pq"
)

type YoutubeVideo struct {
	ID            int            `gorm:"column:id;primaryKey;autoIncrement"`
	ChannelID     int            `gorm:"column:channel_id;not null"`
	Channel       YoutubeChannel `gorm:"foreignKey:ChannelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MinioID       string         `gorm:"column:minio_id"`
	URL           string         `gorm:"column:url"`
	VideoSections pq.StringArray `gorm:"column:video_sections;type:text[]"`
	VideoLength   string         `gorm:"column:video_length"` // Representing SQL INTERVAL as a string
	CreatedAt     time.Time      `gorm:"column:created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at"`
}
