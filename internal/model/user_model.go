package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
