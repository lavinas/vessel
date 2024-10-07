package dto

import (
	"fmt"

	"encoding/json"
)

// CreateClassRequest represents the create class request
type ClassCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Action returns the action of the request
func (r *ClassCreateRequest) Action() string {
	return "ClassCreate"
}

// ToJson returns the json representation of the request
func (r *ClassCreateRequest) ToJson() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// CreateClassResponse represents the create class response
type ClassCreateResponse struct {
	BaseResponse
	ID int64 `json:"id"`
}

// NewClassCreateResponse creates a new class create response
func NewClassCreateResponse(status, description string, ID int64) *ClassCreateResponse {
	return &ClassCreateResponse{
		BaseResponse: BaseResponse{
			Status:      status,
			Description: description,
		},
		ID: ID,
	}
}

// ToJson returns the json representation of the response
func (r *ClassCreateResponse) ToJson() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// ToLine returns the line representation of the response
func (r *ClassCreateResponse) ToLine() string {
	return fmt.Sprintf("%d %s", r.ID, r.Status)
}
