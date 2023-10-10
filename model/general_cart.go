package model

import (
	"gorm.io/gorm"
)

// General Cart struct
type GeneralCart struct {
	gorm.Model
	// IdGeneralCart int     `json:"id_cart" gorm:"primaryKey autoIncrement"`
	Cash      string  `json:"cash"`
	Debit     int     `json:"debit"`
	Credit    int     `json:"credit"`
	IdAccount int     `json:"id_account"`
	Account   Account `gorm:"foreignKey:IdAccount" json:"account"`
}

// General Carts struct
type GeneralCarts struct {
	GeneralCarts []GeneralCart `json:"carts"`
}
