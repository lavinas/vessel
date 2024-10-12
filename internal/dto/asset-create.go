package dto

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	ErrAssetCreateRequestClassNameIsBlank = "class name is blank"
	ErrAssetCreateRequestNameIsBlank      = "name is blank"
	ErrAssetCreateRequestDuplicated = "asset already exists with the same name"
	ErrAssetCreateRequestClassNotFound = "class not found"
)

// AssetCreateRequest represents the create asset request
type AssetCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ClassName   string `json:"class_name"`
}

// Validate validates request params
func (a *AssetCreateRequest) Validate() error {
	if a.ClassName == "" {
		return errors.New(ErrAssetCreateRequestClassNameIsBlank)
	}
	if a.Name == "" {
		return errors.New(ErrAssetCreateRequestNameIsBlank)
	}
	return nil
}

// Action returns the action of the request
func (a *AssetCreateRequest) Action() string {
	return "AssetCreate"
}

// ToJson returns the json representation of the request
func (a *AssetCreateRequest) ToJson() string {
	b, _ := json.Marshal(a)
	return string(b)
}

// AssetCreateResponse represents the create asset response
type AssetCreateResponse struct {
	BaseResponse
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ClassName   string `json:"class_name"`
	CreatedAt   string `json:"created_at"`
}

// NewAssetCreateResponse creates a new asset create response
func NewAssetCreateResponse(status, statusdesc string, ID int64, name, description, className string, createdAt string) *AssetCreateResponse {
	return &AssetCreateResponse{
		BaseResponse: BaseResponse{
			Status:      status,
			Description: statusdesc,
		},
		ID:          ID,
		Name:        name,
		Description: description,
		ClassName:   className,
		CreatedAt:   createdAt,
	}
}

// ToJson returns the json representation of the response
func (a *AssetCreateResponse) ToJson() string {
	b, _ := json.Marshal(a)
	return string(b)
}

// String returns the string representation of the response
func (a *AssetCreateResponse) String() string {
	if a.Status == StatusSuccess {
		ret := [][]string{
			{"ID", "Name", "Description", "Class Name", "Created At"},
			{fmt.Sprintf("%d", a.ID), a.Name, a.Description, a.ClassName, a.CreatedAt},
		}
		return a.ToTable(ret)
	}
	return a.BaseResponse.String()
}
