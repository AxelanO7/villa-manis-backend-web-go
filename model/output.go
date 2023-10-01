package model

import (
	"gorm.io/gorm"
)

// Output struct
type Output struct {
	gorm.Model
	IdOutput     int    `json:"id_output"`
	NoOutput     string `json:"no_output"`
	DateOutput   string `json:"date_output"`
	StatusOutput string `json:"status_output"`
}

// Outputs struct
type Outputs struct {
	Outputs []Output `json:"outputs"`
}
