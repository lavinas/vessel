package domain

import (
	"time"
	"errors"

	"github.com/lavinas/vessel/internal/port"
)

// Event represents the event domain model
type Event struct {
	Base
	ID          int64
	Name        string
	Description string
	CreatedAt   time.Time
}

// GetByID is a method that gets an event by ID
func (e *Event) GetByID(id int64) error {
	sql := `SELECT id, name, description, created_at FROM event WHERE id = $1`
	row := e.Repo.QueryRow(sql, id)
	err := e.Repo.Scan(row, &e.ID, &e.Name, &e.Description, &e.CreatedAt)
	if err != nil {
		return err
	}
	if e.ID == 0 {
		return errors.New(port.ErrEventNotFound)
	}
	return nil
}

// GetByName is a method that gets an event by name
func (e *Event) GetByName(name string) error {
	sql := `SELECT id, name, description, created_at FROM event WHERE name = $1`
	row := e.Repo.QueryRow(sql, name)
	err := e.Repo.Scan(row, &e.ID, &e.Name, &e.Description, &e.CreatedAt)
	if err != nil {
		return err
	}
	if e.ID == 0 {
		return errors.New(port.ErrEventNotFound)
	}
	return nil
}

// Create is a method that creates an event
func (e *Event) Create(name, description string) error {
	sql := `INSERT INTO events (name, description, created_at) VALUES ($1, $2, $3) RETURNING id`
	row := e.Repo.QueryRow(sql, name, description, time.Now())
	err := e.Repo.Scan(row, &e.ID)
	if err != nil {
		return err
	}
	return nil
}
