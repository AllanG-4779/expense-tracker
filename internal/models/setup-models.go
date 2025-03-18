package models

type Category struct {
	ID          uint   `db:"id"`
	Name        string `db:"name"`
	Icon        string `db:"icon"`
	Type        string `db:"type"`
	Description string `db:"description"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
	Deleted     bool   `db:"deleted"`
}

type Transaction struct {
	ID          uint    `db:"id"`
	Amount      float64 `db:"amount"`
	Description string  `db:"description"`
	Date        string  `db:"date"`
	Type        string  `db:"type"`
	CategoryID  uint    `db:"category_id"`
	AccountID   uint    `db:"account_id"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
	Deleted     bool    `db:"deleted"`
}
type Account struct {
	ID        uint    `db:"id"`
	UserID    uint    `db:"user_id"`
	Name      string  `db:"name"`
	Balance   float64 `db:"balance"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
	Deleted   bool    `db:"deleted"`
}

type Budget struct {
	ID         uint    `db:"id"`
	Amount     float64 `db:"amount"`
	Balance    float64 `db:"balance"`
	CategoryID uint    `db:"category_id"`
	UserID     uint    `db:"user_id"`
	StartDate  string  `db:"start_date"`
	EndDate    string  `db:"end_date"`
	CreatedAt  string  `db:"created_at"`
	UpdatedAt  string  `db:"updated_at"`
	Page       int     `db:"page"`
	Size       int     `db:"size"`
	Deleted    bool    `db:"deleted"`
}
