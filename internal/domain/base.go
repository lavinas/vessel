package domain

import (
	"github.com/lavinas/vessel/internal/port"
)

// Base represents the base domain model
type Base struct {
	Repo port.Repository
}
