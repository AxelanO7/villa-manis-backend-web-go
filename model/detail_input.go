package model

import (
	"gorm.io/gorm"
)

// DetailInput struct
type DetailInput struct {
	gorm.Model
	InputInformation string  `json:"input_information"`
	Quantity         int     `json:"quantity"`
	TotalPrice       int     `json:"total_price"`
	StatusCart       int     `json:"status_cart"`
	InputDate        string  `json:"input_date"`
	IdInput          int     `json:"id_input"`
	Input            Input   `gorm:"foreignKey:IdInput" json:"input"`
	IdAccount        int     `json:"id_account"`
	Account          Account `gorm:"foreignKey:IdAccount" json:"account"`
}

// DetailInputs struct
type DetailInputs struct {
	DetailInputs []DetailInput `json:"detail_inputs"`
}
