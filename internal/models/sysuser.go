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
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	Email      string    `db:"email"`
	Username   string    `db:"username"`
	Password   string    `db:"password"`
	ID         uint      `db:"id"`
	Status     string    `db:"status"`
	ProfileUrl string    `db:"profile_url"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Deleted    bool      `db:"deleted"`
}
