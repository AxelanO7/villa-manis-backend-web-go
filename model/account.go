package model

import (
	"gorm.io/gorm"
)

// Account struct
type Account struct {
	gorm.Model
	IdAccount   int      `json:"id_account"`
	NameAccount string   `json:"name_account"`
	Type        string   `json:"type"`
	IdCategory  Category `gorm:"foreignKey:IdCategory"`
}

// Accounts struct
type Accounts struct {
	Accounts []Account `json:"accounts"`
}
