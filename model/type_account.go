package model

import (
	"gorm.io/gorm"
)

// TypeAccount struct
type TypeAccount struct {
	gorm.Model
	TypeAccount []TypeAccount `json:"accounts"`
	// IdDetailInput  int          `json:"id_detail_input"`
	// DetailInput    DetailInput  `gorm:"foreignKey:IdDetailInput" json:"detail_input"`
	// IdDetailOutput int          `json:"id_detail_output"`
	// DetailOutput   DetailOutput `gorm:"foreignKey:IdDetailOutput" json:"detail_output"`
}

// TypeAccounts struct
type TypeAccounts struct {
	TypeAccount []TypeAccount `json:"accounts"`
}
