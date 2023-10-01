package model

import (
	"gorm.io/gorm"
)

// General Cart struct
type GeneralCart struct {
	gorm.Model
	IdGeneralCart int     `json:"id_cart"`
	IdAccount     Account `gorm:"foreignKey:IdAccount"`
	Cash          string  `json:"cash"`
	Debit         string  `json:"debit"`
	Credit        string  `json:"credit"`
}

// General Carts struct
type GeneralCarts struct {
	GeneralCarts []GeneralCart `json:"carts"`
}
