package model

import (
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Name     string `json:"name_user"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Level    int    `json:"level"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}
