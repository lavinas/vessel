package entity

import (
	"github.com/lavinas/vessel/internal/port"
)

const (
	baseName = "assets"
)

// Base represents the base of domain model
type Base struct {
	Repo port.Repository
}

// checkTx is a method that checks if a transaction is nil and initializes it if necessary
func (a *Base) CheckTx(tx interface{}) (interface{}, error) {
	if tx == nil {
		return a.Repo.Begin(baseName)
	}
	return tx, nil
}
