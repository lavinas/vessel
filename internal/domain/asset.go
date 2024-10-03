package domain

import (
	"errors"
	"time"

	"github.com/lavinas/vessel/internal/port"
)

// Asset represents the asset domain model
type Asset struct {
	Base
	ID          int64
	Class       *Class
	Name        string
	Description string
	CreatedAt   time.Time
}

// Create is a method that creates an asset
func (a *Asset) Create(className, name, description string) error {
	a.Class = &Class{}
	if err := a.Class.GetByName(className); err != nil {
		return err
	}
	sql := `INSERT INTO assets (class_id, name, description, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	row := a.Repo.QueryRow(sql, a.Class.ID, name, description, time.Now())
	err := a.Repo.Scan(row, &a.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetByName is a method that gets an asset by name
func (a *Asset) GetByName(name string) error {
	sql := `SELECT id, class_id, name, description, created_at FROM assets WHERE name = $1`
	row := a.Repo.QueryRow(sql, name)
	id := int64(0)
	err := a.Repo.Scan(row, &id, &a.Class.ID, &a.Name, &a.Description, &a.CreatedAt)
	if err != nil {
		return err
	}
	if id == 0 {
		return errors.New(port.ErrAssetNotFound)
	}
	a.Class = &Class{}
	if err := a.Class.GetByID(a.Class.ID); err != nil {
		return err
	}
	return nil
}