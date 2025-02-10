package models

// Meme - модель мема
type Meme struct {
	BaseModel
	MemePic     uint   `gorm:"not null" json:"meme_pic"`
	MemeTopic   uint   `gorm:"not null" json:"meme_topic"`
	TextContent string `gorm:"type:jsonb;not null" json:"text_content"`
	PostLink    string `gorm:"type:text;not null" json:"post_link"`
	Karma       int    `gorm:"not null" json:"karma"`
	Comments    int    `gorm:"not null" json:"comments"`
	Creator     *uint  `gorm:"index" json:"creator"`
}

// MemeKarma - модель системы рейтингов для мемов
type MemeKarma struct {
	BaseModel
	MemeID uint `gorm:"not null;index" json:"meme_id"`
	UserID uint `gorm:"not null;index" json:"user_id"`
	Karma  int  `gorm:"not null" json:"karma"`
}

// MemeComment - комментарии к мемам
type MemeComment struct {
	BaseModel
	MemeID  uint   `gorm:"not null;index" json:"meme_id"`
	UserID  *uint  `gorm:"index" json:"user_id"`
	Comment string `gorm:"type:text;not null" json:"comment"`
}
