package dto

import (
	"fmt"

	"encoding/json"
)

const (
	StatusSuccess             = "success"
	StatusCreated             = "created"
	StatusBadRequest          = "bad request"
	StatusUnauthorized        = "unauthorized"
	StatusForbidden           = "forbidden"
	StatusInternalServerError = "internal server error"
	StatusNotFound            = "not found"
	StatusConflict            = "conflict"
)

// BaseResponse represents the base response
type BaseResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

// NewBaseResponse creates a new base response
func NewBaseResponse(status, description string) *BaseResponse {
	return &BaseResponse{
		Status:      status,
		Description: description,
	}
}

// ToJson returns the json representation of the response
func (r *BaseResponse) ToJson() string {
	b, _ := json.Marshal(r)
	return string(b)	
}

// ToLine returns the line representation of the response
func (r *BaseResponse) ToLine() string {
	return fmt.Sprintf("%s %s", r.Status, r.Description)
}