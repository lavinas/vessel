package dto

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