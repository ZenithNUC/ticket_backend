package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(25);not null",json:"name"`
	Password string `gorm:"type:varchar(255);not null",json:"password"`
	Phone    string `gorm:"type:varchar(255);not null",json:"phone"`
}
