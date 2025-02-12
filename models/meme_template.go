package models

import (
	"gorm.io/datatypes"
)

type MemeTemplate struct {
	ID             uint           `gorm:"primaryKey"`
	PicPath        string         `gorm:"column:pic_path;type:text;not null"`
	AdditionalInfo datatypes.JSON `gorm:"column:additional_info;type:jsonb"`
	TextContent    datatypes.JSON `gorm:"column:text_content;type:jsonb"`
	IsProccessing  bool           `gorm:"column:is_proccessing;default:false"`
	// Один шаблон может использоваться в нескольких сгенерированных мемах
	GeneratedMemes []GeneratedMeme `gorm:"foreignKey:MemeTemplateID"`
}
