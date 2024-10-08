package entity

import (
	"strconv"
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

// GetByID is a method that gets a class by ID
func (c *Class) GetByID(id int64, tx interface{}) error {
	tx, err := c.CheckTx(tx)
	if err != nil {
		return err
	}
	fields := []string{"id", "name", "description", "created_at"}
	vals, err := c.Repo.GetId(tx, baseName, classTable, id, &fields)
	if err != nil {
		return err
	}
	if c.ID, err = strconv.ParseInt((*vals)[0], 10, 64); err != nil {
		return err
	}
	c.Name = (*vals)[1]
	c.Description = (*vals)[2]
	if c.CreatedAt, err = time.Parse(time.DateTime, (*vals)[3]); err != nil {
		return err
	}
	return nil
}

// Get is a method that gets a class
func (c *Class) GetByName(name string, tx interface{}) error {
	fields := []string{"id", "name", "description", "created_at"}
	vals, err := c.Repo.GetField(tx, baseName, classTable, "name", name, &fields)
	if err != nil {
		return err
	}
	if c.ID, err = strconv.ParseInt((*vals)[0], 10, 64); err != nil {
		return err
	}
	c.Name = (*vals)[1]
	c.Description = (*vals)[2]
	if c.CreatedAt, err = time.Parse(time.DateTime, (*vals)[3]); err != nil {
		return err
	}
	return nil
}

// Create is a method that creates a class
func (c *Class) Create(name, description string, tx interface{}) error {
	var err error
	fds := []string{"name", "description", "created_at"}
	vals := []string{name, description, time.Now().Format(time.DateTime)}
	if c.ID, err = c.Repo.InsertAuto(tx, baseName, classTable, &fds, &vals); err != nil {
		return err
	}
	return nil
}
