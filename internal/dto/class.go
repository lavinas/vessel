package dto

// CreateClassRequest represents the create class request
type CreateClassRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateClassResponse represents the create class response
type CreateClassResponse struct {
	ID          int64  `json:"id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

// GetClassRequest represents the get class request
type GetClassRequest struct {
	ID int64 `json:"id"`
}

// GetClassResponse represents the get class response
type GetClassResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
