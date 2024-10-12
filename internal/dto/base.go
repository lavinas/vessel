package dto

import (
	"fmt"
	"strings"

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

	ErrInternalGeneric = "contact administrator"
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
func (r *BaseResponse) String() string {
	return fmt.Sprintf("%s: %s", r.Status, r.Description)
}

// mountTable is a function that mounts a table
func (c *BaseResponse) ToTable(table [][]string) string {
	ret := ""
	// get number of columns from the first table row
	columnLengths := make([]int, len(table[0]))
	for _, line := range table {
		for i, val := range line {
			if len(val) > columnLengths[i] {
				columnLengths[i] = len(val)
			}
		}
	}
	var lineLength int
	for _, c := range columnLengths {
		lineLength += c + 3
	}
	lineLength += 1
	for i, line := range table {
		if i == 0 {
			ret += fmt.Sprintf("+%s+\n", strings.Repeat("-", lineLength-2))
		}
		for j, val := range line {
			ret += fmt.Sprintf("| %-*s ", columnLengths[j], val)
			if j == len(line)-1 {
				ret += "|\n"
			}
		}
		if i == 0 || i == len(table)-1 {
			ret += fmt.Sprintf("+%s+\n", strings.Repeat("-", lineLength-2))
		}
	}
	return ret
}
