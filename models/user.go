package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"column:username;type:text;not null"`
	PicPath   string    `gorm:"column:pic_path;type:text"`
	Password  string    `gorm:"column:password;type:text;not null"`
	Karma     int       `gorm:"column:karma"`
	Userpic   string    `gorm:"column:userpic;type:text"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	// Один пользователь может создавать несколько наших Reddit-постов
	OurRedditPosts []OurRedditPost `gorm:"foreignKey:CreatorID"`
}
