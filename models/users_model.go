package models

type User struct {
	ID    int    `json:"id" form:"id" gorm:"primaryKey"`
	Name  string `json:"name" form:"name" gorm:"not null"`
	Email string `json:"email" form:"email" gorm:"not null"`
}
