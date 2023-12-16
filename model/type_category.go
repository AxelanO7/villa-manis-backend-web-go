package model

import (
	"gorm.io/gorm"
)

// TypeCategory struct
type TypeCategory struct {
	gorm.Model
	IdAccount int     `json:"id_account"`
	Account   Account `gorm:"foreignKey:IdAccount" json:"account"`
}

// TypeCategorys struct
type TypeCategorys struct {
	TypeCategory []TypeCategory `json:"categorys"`
}
