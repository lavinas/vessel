package entity

import (
	"time"
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
func (e *Event) GetByID(id int64, tx interface{}) error {
	tx, err := e.CheckTx(tx)
	if err != nil {
		return err
	}
	fields := []string{"id", "name", "description", "created_at"}
	vals, err := e.Repo.GetId(tx, baseName, "event", id, &fields)
	if err != nil {
		return err
	}
	e.ID = (*vals)[0].(int64)
	e.Name = (*vals)[1].(string)
	e.Description = (*vals)[2].(string)
	e.CreatedAt = (*vals)[3].(time.Time)
	return nil
}

// GetByName is a method that gets an event by name
func (e *Event) GetByName(name string, tx interface{}) error {
	tx, err := e.CheckTx(tx)
	if err != nil {
		return err
	}
	fields := []string{"id", "name", "description", "created_at"}
	vals, err := e.Repo.GetField(tx, baseName, "event", "name", name, &fields)
	if err != nil {
		return err
	}
	e.ID = (*vals)[0].(int64)
	e.Name = (*vals)[1].(string)
	e.Description = (*vals)[2].(string)
	e.CreatedAt = (*vals)[3].(time.Time)
	return nil
}

// Create is a method that creates an event
func (e *Event) Create(name, description string) error {
	tx, err := e.CheckTx(nil)
	if err != nil {
		return err
	}
	fds := []string{"name", "description", "created_at"}
	vals := []string{name, description, time.Now().Format(time.DateTime)}
	if e.ID, err = e.Repo.InsertAuto(tx, baseName, "event", &fds, &vals); err != nil {
		return err
	}
	return nil
}
