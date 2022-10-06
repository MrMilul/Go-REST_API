package models

type User struct {
	Id       int    `json:"id" gorm:"privateKey"`
	Username string `json:"username" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
