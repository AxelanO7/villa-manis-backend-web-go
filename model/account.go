package model

import (
	"gorm.io/gorm"
)

// Account struct
type Account struct {
	gorm.Model
	Code        string   `json:"code"`
	NameAccount string   `json:"name_account"`
	Character   string   `json:"character"`
	IdCategory  int      `json:"id_category"`
	Category    Category `gorm:"foreignKey:IdCategory" json:"category"`
}

// Accounts struct
type Accounts struct {
	Accounts []Account `json:"accounts"`
}
