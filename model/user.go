package model

import (
	// "github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	// IdUser       uuid.UUID `gorm:"type:uuid;"`
	IdUser   string `json:"id_user"`
	Name     string `json:"name_user"`
	Username string `json:"username"`
	// Email    string `json:"email"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
	Level    string `json:"level"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}

// func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	// UUID version 4
// 	user.ID = uuid.New()
// 	return
// }
