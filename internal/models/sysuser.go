package models

import (
	"gorm.io/gorm"
)

type SystemUser struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string   `gorm:"uniqueIndex;not null"`
	Username  string   `gorm:"uniqueIndex;not null"`
	Password  string   `gorm:"type:varchar(400);not null"`
	Status    bool     `gorm:"default:false"`
	Account   Account  `gorm:"foreignKey:UserID"`
	Budget    []Budget `gorm:"foreignKey:UserID"`
}
