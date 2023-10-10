package model

import (
	"gorm.io/gorm"
)

// Account struct
type Account struct {
	gorm.Model
	NameAccount string   `json:"name_account"`
	Type        string   `json:"type"`
	IdCategory  int      `json:"id_category"`
	Category    Category `gorm:"foreignKey:IdCategory" json:"category"`
}

// Accounts struct
type Accounts struct {
	Accounts []Account `json:"accounts"`
}
