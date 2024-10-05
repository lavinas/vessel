package dto

import (
	"encoding/json"
)

// GetClassRequest represents the get class request
type ClassGetRequest struct {
	ID int64 `json:"id"`
}

// Action returns the action of the request
func (r *ClassGetRequest) Action() string {
	return "ClassGet"
}

// ToJson returns the json representation of the request
func (r *ClassGetRequest) ToJson() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// GetClassResponse represents the get class response
type ClassGetResponse struct {
	BaseResponse
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

// NewClassGetResponse creates a new class get response
func NewClassGetResponse(status, statusdesc string, ID int64, name, description string, createdAt string) *ClassGetResponse {
	return &ClassGetResponse{
		BaseResponse: BaseResponse{
			Status:      status,
			Description: statusdesc,
		},
		ID:          ID,
		Name:        name,
		Description: description,
		CreatedAt:   createdAt,
	}
}
