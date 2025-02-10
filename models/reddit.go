package models

// RedditPost - модель поста с Reddit
type RedditPost struct {
	BaseModel
	PicPath     string `gorm:"type:text;not null" json:"pic_path"`
	PostLink    string `gorm:"type:text;not null" json:"post_link"`
	TreadName   string `gorm:"type:text;not null" json:"tread_name"`
	TextContent string `gorm:"type:jsonb;not null" json:"text_content"`
	TreadLink   string `gorm:"type:text;not null" json:"tread_link"`
	PostTitle   string `gorm:"type:text;not null" json:"post_title"`
	PostText    string `gorm:"type:text;not null" json:"post_text"`
	Karma       int    `gorm:"not null" json:"karma"`
	Comments    int    `gorm:"not null" json:"comments"`
}
