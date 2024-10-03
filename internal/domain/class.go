package domain

import (
	"time"
	"errors"

	"github.com/lavinas/vessel/internal/port"
)

// Class represents the class domain model
type Class struct {
	Base
	ID          int64
	Name        string
	Description string
	CreatedAt   time.Time
}


// GetByID is a method that gets a class by ID
func (c *Class) GetByID(id int64) error {
	sql := `SELECT id, name, description, created_at FROM classes WHERE id = $1`
	row := c.Repo.QueryRow(sql, id)
	err := c.Repo.Scan(row, &c.ID, &c.Name, &c.Description, &c.CreatedAt)
	if err != nil {
		return err
	}
	if c.ID == 0 {
		return errors.New(port.ErrClassNotFound)
	}
	return nil
}

// Get is a method that gets a class
func (c *Class) GetByName(name string) error {
	sql := `SELECT id, name, description, created_at FROM classes WHERE nam = '$1'`
	row := c.Repo.QueryRow(sql, name)
	err := c.Repo.Scan(row, &c.ID, &c.Name, &c.Description, &c.CreatedAt)
	if err != nil {
		return err
	}
	if c.ID == 0 {
		return errors.New(port.ErrClassNotFound)
	}
	return nil
}

// Create is a method that creates a class
func (c *Class) Create(name, description string) error {
	sql := `INSERT INTO classes (name, description, created_at) VALUES ($1, $2, $3) RETURNING id`
	row := c.Repo.QueryRow(sql, name, description, time.Now())
	err := c.Repo.Scan(row, &c.ID)
	if err != nil {
		return err
	}
	return nil
}

