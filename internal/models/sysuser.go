package models

import (
	"time"

	"gorm.io/gorm"
)

type LoginAccount struct {
	gorm.Model
	Username   string `gorm:"unique"`
	Password   string `gorm:"size:256;not-null"`
	FirstLogin bool   `gorm:"default:true"`
	LastLogin  time.Time
	LoginToken int
	UserID uint 
}

type SystemUser struct {
	gorm.Model
	FirstName   string `gorm:"type:varchar(40)"`
	LastName    string `gorm:"type:varchar(40)"`
	Email       string`gorm:"type:varchar(100)"`
	Dob         string 
	Residential string
	Phone       string
	LoginAccount      LoginAccount `gorm:"foreignKey:UserID"`
}
