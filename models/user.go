package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	First_name string `json:"first_name" gorm:"not null"`
	Last_name  string `json:"last_name" gorm:"not null"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password" gorm:"size:32 not null"`
}
