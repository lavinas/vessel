package dto

const (
	StatusSuccess  = "success"
	StatusCreated  = "created"
	StatusFailed   = "failed"
	StatusNotFound = "not found"
	StatusConflict = "conflict"
)

// BaseResponse represents the base response
type BaseResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}
