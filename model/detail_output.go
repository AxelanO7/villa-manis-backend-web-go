package model

import (
	"gorm.io/gorm"
)

// DetailOutput struct
type DetailOutput struct {
	gorm.Model
	IdDetailOutput    int    `json:"id_detail_output"`
	NoOutput          string `json:"no_output"`
	IdCash            int    `json:"id_cash"`
	IdAccount         int    `json:"id_account"`
	OutputInformation string `json:"output_information"`
	Quantity          int    `json:"quantity"`
	TotalPrice        int    `json:"total_price"`
	StatusCart        string `json:"status_cart"`
	OutputDate        string `json:"output_date"`
}

// DetailOutputs struct
type DetailOutputs struct {
	DetailOutputs []DetailOutput `json:"detail_outputs"`
}
