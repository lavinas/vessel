package domain

import (
	"fmt"
	"strconv"
	"time"
)

const (
	assetTable = "asset"
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
func (a *Asset) Create(className, name, description string, tx interface{}) error {
	tx, err := a.CheckTx(tx)
	if err != nil {
		return err
	}
	a.Class = &Class{}
	if err := a.Class.GetByName(className, tx); err != nil {
		return err
	}
	fds := []string{"class_id", "name", "description", "created_at"}
	vals := []string{fmt.Sprintf("%d", a.Class.ID), name, description, time.Now().Format(time.DateTime)}
	if a.ID, err = a.Repo.InsertAuto(tx, base, assetTable, &fds, &vals); err != nil {
		return err
	}	
	return nil
}

// GetByName is a method that gets an asset by name
func (a *Asset) GetByName(name string, tx interface{}) error {
	tx, err := a.CheckTx(tx)
	if err != nil {
		return err
	}
	fields := []string{"id", "class_id", "name", "description", "created_at"}
	vals, err := a.Repo.GetId(tx, base, assetTable, a.ID, &fields)
	if err != nil {
		return err
	}
	if a.ID, err = strconv.ParseInt((*vals)[0], 10, 64); err != nil {
		return err
	}
	a.Class = &Class{}
	if a.Class.ID, err = strconv.ParseInt((*vals)[1], 10, 64); err != nil {
		return err
	}
	if err := a.Class.GetByID(a.Class.ID, tx); err != nil {
		return err
	}
	a.Name = (*vals)[2]
	a.Description = (*vals)[3]
	a.CreatedAt, err = time.Parse(time.DateTime, (*vals)[4])
	if err != nil {
		return err
	}
	return nil
}
