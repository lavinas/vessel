package entity

import (
	"time"

	"github.com/lavinas/vessel/internal/port"
)

const (
	ErrNameIsBlank = "name is blank"
	classTable     = "class"
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
	tx, err := c.CheckTx(tx)
	if err != nil {
		return err
	}
	where := map[string]interface{}{"id": id}	
	vals, err := c.Repo.Get(tx, baseName, classTable, &where)
	if err != nil {
		return err
	}
	if len(*vals) == 0 {
		return nil
	}
	c.ID = (*vals)[0]["id"].(int64)
	c.Name = (*vals)[0]["name"].(string)
	c.Description = (*vals)[0]["description"].(string)
	c.CreatedAt = (*vals)[0]["created_at"].(time.Time)
	return nil
}

// Get is a method that gets a class
func (c *Class) GetByName(name string, tx interface{}) error {
	where := map[string]interface{}{"name": name}
	vals, err := c.Repo.Get(tx, baseName, classTable, &where)
	if err != nil {
		return err
	}
	if len(*vals) == 0 {
		return nil
	}
	c.ID = (*vals)[0]["id"].(int64)
	c.Name = (*vals)[0]["name"].(string)
	c.Description = (*vals)[0]["description"].(string)
	c.CreatedAt = (*vals)[0]["created_at"].(time.Time)
	return nil
}


