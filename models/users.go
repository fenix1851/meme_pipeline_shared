package models

// User - модель пользователя
type User struct {
	BaseModel
	Username string `gorm:"type:text;unique;not null" json:"username"`
	Password string `gorm:"type:text;not null" json:"password"`
	Karma    int    `gorm:"not null" json:"karma"`
	Userpic  string `gorm:"type:text" json:"userpic"`
}
