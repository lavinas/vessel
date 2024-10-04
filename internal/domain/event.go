package domain

import (
	"time"
	"strconv"
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
	vals, err := e.Repo.GetId(tx, base, "event", id, &fields)
	if err != nil {
		return err
	}
	if e.ID, err = strconv.ParseInt((*vals)[0], 10, 64); err != nil {
		return err
	}
	e.Name = (*vals)[1]
	e.Description = (*vals)[2]
	if e.CreatedAt, err = time.Parse(time.DateTime, (*vals)[3]); err != nil {
		return err
	}
	return nil
}

// GetByName is a method that gets an event by name
func (e *Event) GetByName(name string, tx interface{}) error {
	tx, err := e.CheckTx(tx)
	if err != nil {
		return err
	}
	fields := []string{"id", "name", "description", "created_at"}
	vals, err := e.Repo.GetField(tx, base, "event", "name", name, &fields)
	if err != nil {
		return err
	}
	if e.ID, err = strconv.ParseInt((*vals)[0], 10, 64); err != nil {
		return err
	}
	e.Name = (*vals)[1]
	e.Description = (*vals)[2]
	if e.CreatedAt, err = time.Parse(time.DateTime, (*vals)[3]); err != nil {
		return err
	}
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
	if e.ID, err = e.Repo.InsertAuto(tx, base, "event", &fds, &vals); err != nil {
		return err
	}
	return nil
}
