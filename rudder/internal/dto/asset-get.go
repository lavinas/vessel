package dto

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	ErrAssetGetRequestInvalidID = "invalid ID"
	ErrAssetGetRequestNotFound  = "asset not found"
)

// AssetGetRequest represents the get asset request
type AssetGetRequest struct {
	ID int64 `json:"id"`
}

// Validate validates request params
func (r *AssetGetRequest) Validate() error {
	if r.ID == 0 {
		return errors.New(ErrAssetGetRequestInvalidID)
	}
	return nil
}

// Action returns the action of the request
func (r *AssetGetRequest) Action() string {
	return "AssetGet"
}

// ToJson returns the json representation of the request
func (r *AssetGetRequest) ToJson() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// AssetGetResponse represents the get asset response
type AssetGetResponse struct {
	BaseResponse
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ClassName   string `json:"class_name"`
	CreatedAt   string `json:"created_at"`
}

// NewAssetGetResponse creates a new asset get response
func NewAssetGetResponse(status, statusdesc string, ID int64, name, description, className string, createdAt string) *AssetGetResponse {
	return &AssetGetResponse{
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
func (a *AssetGetResponse) ToJson() string {
	b, _ := json.Marshal(a)
	return string(b)
}

// String returns the string representation of the response
func (a *AssetGetResponse) String() string {
	if a.Status == StatusSuccess {
		ret := [][]string{
			{"ID", "Name", "Description", "Class Name", "Created At"},
			{fmt.Sprintf("%d", a.ID), a.Name, a.Description, a.ClassName, a.CreatedAt},
		}
		return a.ToTable(ret)
	}
	return a.BaseResponse.String()
}
