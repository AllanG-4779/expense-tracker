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
	Type        string  `json:"type"`
	CategoryID  uint    `json:"category_id"`
	AccountID   uint    `json:"account_id"`
}
