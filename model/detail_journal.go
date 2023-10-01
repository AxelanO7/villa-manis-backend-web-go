package model

import (
	"gorm.io/gorm"
)

// General Cart struct
type DetailJournal struct {
	gorm.Model
	IdDetailJournal    int            `json:"id_detail_journal"`
	IdGeneralJournal   GeneralJournal `gorm:"foreignKey:IdGeneralJournal"`
	JournalInformation string         `json:"journal_information"`
	RefJournal         string         `json:"ref_journal"`
	IdAccount          int            `json:"id_account"`
	Debit              string         `json:"debit"`
	Credit             string         `json:"credit"`
	DateTransaction    string         `json:"date_transaction"`
	StatusPost         string         `json:"status_post"`
}

// General Carts struct
type DetailJournals struct {
	DetailJournals []DetailJournal `json:"detail_journals"`
}
