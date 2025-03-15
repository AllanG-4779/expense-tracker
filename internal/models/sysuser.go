package models

import (
	"time"
)

type LoginAccount struct {
	Username   string
	Password   string
	FirstLogin bool
	LastLogin  time.Time
	LoginToken int
	UserID     uint
}

type SystemUser struct {
	FirstName    string
	LastName     string
	Email        string
	Dob          string
	Residential  string
	Phone        string
	LoginAccount LoginAccount
}
