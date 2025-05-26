package types

type CategoryRequest struct {
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Page        int    `json:"page"`
	Size        int    `json:"size"`
}
type AccountRequest struct {
	Name    string  `json:"name"`
	Balance float32 `json:"balance"`
	UserId  uint    `json:"user_id"`
	Page    int     `json:"page"`
	Size    int     `json:"size"`
}
type TransactionRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Title       string  `json:"title"`
	Type        string  `json:"type"`
	CategoryID  string  `json:"category"`
	AccountID   uint    `json:"account_id"`
	CategoryId uint    `json:"category_id"`
	TransactionID uint   `json:"transaction_id"`
	UserId      uint    `json:"user_id"`
	Page        int     `json:"page"`
	Size        int     `json:"size"`
}
type BudgetRequest struct {
	Amount       float64 `json:"amount"`
	Balance      float64 `json:"balance"`
	CategoryName string  `json:"category"`
	StartDate    string  `json:"start_date"`
	UserId       uint    `json:"user_id"`
	EndDate      string  `json:"end_date"`
}
type FetchRequest struct {
	Page     int    `json:"page"`
	ID	   int   `json:"id"`
	Username string `json:"username"`
	Size     int    `json:"size"`
}
