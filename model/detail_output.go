package model

import (
	"gorm.io/gorm"
)

// DetailOutput struct
type DetailOutput struct {
	gorm.Model
	IdCash            int     `json:"id_cash"`
	OutputInformation string  `json:"output_information"`
	Quantity          int     `json:"quantity"`
	TotalPrice        int     `json:"total_price"`
	StatusCart        int     `json:"status_cart"`
	OutputDate        string  `json:"output_date"`
	IdAccount         int     `json:"id_account"`
	Account           Account `gorm:"foreignKey:IdAccount" json:"account"`
	IdOutput          int     `json:"id_output"`
	Output            Output  `gorm:"foreignKey:IdOutput" json:"output"`
}

// DetailOutputs struct
type DetailOutputs struct {
	DetailOutputs []DetailOutput `json:"detail_outputs"`
}
