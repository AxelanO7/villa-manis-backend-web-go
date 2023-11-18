package model

import (
	"gorm.io/gorm"
)

// Input struct
type Input struct {
	gorm.Model
	NoInput     string `json:"no_input"`
	DateInput   string `json:"date_input"`
	StatusInput string `json:"status_input"`
}

// Inputs struct
type Inputs struct {
	Inputs []Input `json:"inputs"`
}
