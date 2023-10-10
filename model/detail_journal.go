package model

import (
	"gorm.io/gorm"
)

// General Cart struct
type DetailJournal struct {
	gorm.Model
	JournalInformation string         `json:"journal_information"`
	RefJournal         string         `json:"ref_journal"`
	Debit              int            `json:"debit"`
	Credit             int            `json:"credit"`
	DateTransaction    string         `json:"date_transaction"`
	StatusPost         int            `json:"status_post"`
	IdGeneralJournal   int            `json:"id_general_journal"`
	GeneralJournal     GeneralJournal `gorm:"foreignKey:IdGeneralJournal" json:"general_journal"`
	IdAccount          int            `json:"id_account"`
	Account            Account        `gorm:"foreignKey:IdAccount" json:"account"`
}

// General Carts struct
type DetailJournals struct {
	DetailJournals []DetailJournal `json:"detail_journals"`
}
