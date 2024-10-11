package dto

import (
	"errors"
	"fmt"
	"strings"

	"encoding/json"
)

const (
	ErrClassCreateRequestNameIsBlank = "name is blank"
)

// CreateClassRequest represents the create class request
type ClassCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Validate validates request params
func (r *ClassCreateRequest) Validate() error {
	if strings.Trim(r.Name, " ") == "" {
		return errors.New(ErrClassCreateRequestNameIsBlank)
	}
	return nil
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
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

// NewClassCreateResponse creates a new class create response
func NewClassCreateResponse(status, statusdesc string, ID int64, name, description string, createdAt string) *ClassCreateResponse {
	return &ClassCreateResponse{
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
func (r *ClassCreateResponse) ToJson() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// ToLine returns the line representation of the response
func (r *ClassCreateResponse) String() string {
	if r.Status == StatusSuccess {
		ret := [][]string{
			{"ID", "Name", "Description", "Created At"},
			{fmt.Sprintf("%d", r.ID), r.Name, r.Description, r.CreatedAt},
		}
		return r.ToTable(ret)
	}
	return fmt.Sprintf("error: %s - %s", r.BaseResponse.Status, r.BaseResponse.Description)
}
