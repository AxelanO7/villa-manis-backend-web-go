package model

import (
	"gorm.io/gorm"
)

// General Journal struct
type GeneralJournal struct {
	gorm.Model
	IdGeneralJournal     int    `json:"id_journal"`
	NoGeneralJournal     string `json:"no_journal"`
	DateGeneralJournal   string `json:"date_journal"`
	StatusGeneralJournal string `json:"status_journal"`
}

// General Journals struct
type GeneralJournals struct {
	GeneralJournals []GeneralJournal `json:"journals"`
}
