package models

import (
	"gorm.io/datatypes"
)

type MemeTemplate struct {
	ID          uint           `gorm:"primaryKey"`
	PicPath     string         `gorm:"column:pic_path;type:text;not null"`
	TextContent datatypes.JSON `gorm:"column:text_content;type:jsonb"`
	// Один шаблон может использоваться в нескольких сгенерированных мемах
	GeneratedMemes []GeneratedMeme `gorm:"foreignKey:MemeTemplateID"`
}
