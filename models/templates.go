package models

// MemeTemplate - шаблон для мема
type MemeTemplate struct {
	BaseModel
	PicPath     string `gorm:"type:text;not null" json:"pic_path"`
	TextContent string `gorm:"type:jsonb;not null" json:"text_content"`
}
