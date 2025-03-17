package types

type CategoryRequest struct {
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Page        int    `json:"page"`
	Size        int    `json:"size"`
}
