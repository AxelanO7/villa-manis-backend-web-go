package model

import (
	"gorm.io/gorm"
)

// Input struct
type Input struct {
	gorm.Model
	IdInput int    `json:"id_input"`
	NoInput string `json:"no_input"`
	Date    string `json:"date_input"`
	// todo : status_input
	Status string `json:"status_input"`
}

// Inputs struct
type Inputs struct {
	Inputs []Input `json:"inputs"`
}
