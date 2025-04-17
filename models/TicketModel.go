package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model  `gorm:"unique"`
	FirstName   string
	LastName    string
	Email       string
	TicketCount uint
}
