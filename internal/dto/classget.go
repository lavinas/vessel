package dto

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	ErrClassGetRequestInvalidID = "invalid ID"
)

// GetClassRequest represents the get class request
type ClassGetRequest struct {
	ID int64 `json:"id"`
}

// Validate validates request params
func (r *ClassGetRequest) Validate() error {
	if r.ID == 0 {
		return errors.New(ErrClassGetRequestInvalidID)
	}
	return nil
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

// ToJson returns the json representation of the response
func (r *ClassGetResponse) ToJson() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// ToLine returns the line representation of the response
func (r *ClassGetResponse) ToLine() string {
	return fmt.Sprintf("%d %s %s %s %s", r.ID, r.Name, r.Description, r.CreatedAt, r.Status)
}
