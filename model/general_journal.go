package model

import (
	"gorm.io/gorm"
)

// General Journal struct
type GeneralJournal struct {
	gorm.Model
	NoGeneralJournal     string `json:"no_journal"`
	DateGeneralJournal   string `json:"date_journal"`
	StatusGeneralJournal int    `json:"status_journal"`
}

// General Journals struct
type GeneralJournals struct {
	GeneralJournals []GeneralJournal `json:"journals"`
}
