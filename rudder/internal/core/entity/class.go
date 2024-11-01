package entity

import (
	"time"

	"github.com/lavinas/vessel/internal/port"
)

const (
	classTable = "class"
)

// Class represents the class domain model
type Class struct {
	Base
	ID          int64
	Name        string
	Description string
	CreatedAt   time.Time
}

// NewClass creates a new Class
func NewClass(repo port.Repository) *Class {
	return &Class{
		Base: Base{
			Repo: repo,
		},
	}
}

// Create is a method that creates a class
func (c *Class) Create(name, description string, tx interface{}) error {
	var err error
	vls := map[string]interface{}{
		"name":        name,
		"description": description,
		"created_at":  time.Now().Format(time.DateTime),
	}
	if c.ID, err = c.Repo.Insert(tx, baseName, classTable, &vls); err != nil {
		return err
	}
	return nil
}

// GetByID is a method that gets a class by ID
func (c *Class) GetByID(id int64, tx interface{}) error {
	return c.get(&map[string]interface{}{"id": id}, tx)
}

// Get is a method that gets a class
func (c *Class) GetByName(name string, tx interface{}) error {
	return c.get(&map[string]interface{}{"name": name}, tx)
}

// Loaded is a method that checks if a class is loaded
func (c *Class) IsLoaded() bool {
	return c.ID != 0
}

// CheckDuplicity is a method that checks if a class is duplicated
func (c *Class) CheckDuplicity(name string, tx interface{}) (bool, error) {
	where := map[string]interface{}{
		"name": name,
	}
	vals, err := c.Repo.Get(tx, baseName, classTable, &where)
	if err != nil {
		return false, err
	}
	return len(*vals) > 0, nil
}

// get is a method that gets a class from the database
func (c *Class) get(where *map[string]interface{}, tx interface{}) error {
	vals, err := c.Repo.Get(tx, baseName, classTable, where)
	if err != nil {
		return err
	}
	if len(*vals) == 0 {
		return nil
	}
	c.ID = (*vals)[0]["id"].(int64)
	c.Name = (*vals)[0]["name"].(string)
	c.Description = (*vals)[0]["description"].(string)
	if c.CreatedAt, err = time.Parse(time.DateTime, (*vals)[0]["created_at"].(string)); err != nil {
		return err
	}
	return nil
}
