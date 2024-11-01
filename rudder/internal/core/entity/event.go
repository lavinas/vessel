package entity

import (
	"time"
)

const (
	eventTable = "event"
)

// Event represents the event domain model
type Event struct {
	Base
	ID          int64
	Name        string
	Description string
	CreatedAt   time.Time
}

// Create is a method that creates an event
func (e *Event) Create(name, description string) error {
	tx, err := e.CheckTx(nil)
	if err != nil {
		return err
	}
	vls := map[string]interface{}{
		"name":        name,
		"description": description,
		"created_at":  time.Now().Format(time.DateTime),
	}
	if e.ID, err = e.Repo.Insert(tx, baseName, eventTable, &vls); err != nil {
		return err
	}
	return nil
}

// GetByID is a method that gets an event by ID
func (e *Event) GetByID(id int64, tx interface{}) error {
	tx, err := e.CheckTx(tx)
	if err != nil {
		return err
	}
	vals, err := e.Repo.Get(tx, baseName, eventTable, &map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if len(*vals) == 0 {
		return nil
	}
	e.ID = (*vals)[0]["id"].(int64)
	e.Name = (*vals)[0]["name"].(string)
	e.Description = (*vals)[0]["description"].(string)
	e.CreatedAt = (*vals)[0]["created_at"].(time.Time)
	return nil
}

// GetByName is a method that gets an event by name
func (e *Event) GetByName(name string, tx interface{}) error {
	tx, err := e.CheckTx(tx)
	if err != nil {
		return err
	}
	vals, err := e.Repo.Get(tx, baseName, eventTable, &map[string]interface{}{"name": name})
	if err != nil {
		return err
	}
	if len(*vals) == 0 {
		return nil
	}
	e.ID = (*vals)[0]["id"].(int64)
	e.Name = (*vals)[0]["name"].(string)
	e.Description = (*vals)[0]["description"].(string)
	e.CreatedAt = (*vals)[0]["created_at"].(time.Time)
	return nil
}
