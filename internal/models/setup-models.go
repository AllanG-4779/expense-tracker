package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name         string        `gorm:"size:255;not null;unique"`
	Icon         string        `gorm:"type:text;not null"`
	Type         string        `gorm:"type:text;check:type in ('expense','income')"`
	Description  string        `gorm:"type:text;not null"`
	Transactions []Transaction `gorm:"foreignkey:CategoryID"`
	Budget       []Budget      `gorm:"foreignkey:CategoryID"`
}
type Account struct {
	gorm.Model
	UserID  uint `gorm:"uniqueIndex"`
	Name    string
	Balance float64
}
type Transaction struct {
	gorm.Model
	Amount      float64 `gorm:"type:decimal(10,2);not null"`
	Description string  `gorm:"type:text"`
	Date        string  `gorm:"type:date;not null"`
	Type        string  `gorm:"check:type in ('expense','income')"` //check ( type in ('expense', 'income') ),
	CategoryID  uint    `gorm:"not null"`
	UserID      uint    `gorm:"null;"`
	AccountID   uint    `gorm:"not null"`
}

type Budget struct {
	gorm.Model
	Amount      float64 `gorm:"not null"`
	Balance     float64 `gorm:"not null;default 0.0"`
	CategoryID  uint    `gorm:"not null"`
	UserID      uint    `gorm:"not null"`
	StartDate   string  `gorm:"start_date"`
	Utilization float64 `gorm:"not null;default 0.0"`
	EndDate     string  `gorm:"end_date"`
	Fraction    float64 `gorm:"not null;default 0.0"`
}
