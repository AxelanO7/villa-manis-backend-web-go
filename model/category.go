package model

import (
	"gorm.io/gorm"
)

// Category struct
type Category struct {
	gorm.Model
	IdCategory   int    `json:"id_category"`
	NameCategory string `json:"name_category"`
}

// Categorys struct
type Categorys struct {
	Categorys []Category `json:"categorys"`
}
