package model

import (
	"gorm.io/gorm"
)

// DetailInput struct
type DetailInput struct {
	gorm.Model
	IdDetailInput    int    `json:"id_detail_input"`
	NoInput          string `json:"no_input"`
	IdAccount        int    `json:"id_account"`
	InputInformation string `json:"input_information"`
	Quantity         int    `json:"quantity"`
	TotalPrice       int    `json:"total_price"`
	StatusCart       string `json:"status_cart"`
	InputDate        string `json:"input_date"`
}

// DetailInputs struct
type DetailInputs struct {
	DetailInputs []DetailInput `json:"detail_inputs"`
}
