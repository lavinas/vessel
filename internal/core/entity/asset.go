package entity

import (
	"fmt"
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
	if a.ID, err = a.Repo.InsertAuto(tx, baseName, assetTable, &fds, &vals); err != nil {
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
	vals, err := a.Repo.GetId(tx, baseName, assetTable, a.ID, &fields)
	if err != nil {
		return err
	}
	a.ID = (*vals)[0].(int64)
	a.Name = (*vals)[2].(string)
	a.Description = (*vals)[3].(string)
	a.CreatedAt = (*vals)[4].(time.Time)
	a.Class, err = a.getAssetClass((*vals)[1].(int64), tx)
	if err != nil {
		return err
	}
	return nil
}

// getAssetClass is a method that gets the asset class
func (a *Asset) getAssetClass(id int64, tx interface{}) (*Class, error) {
	a.Class = &Class{}
	a.Class.ID = id
	if err := a.Class.GetByID(a.Class.ID, tx); err != nil {
		return nil, err
	}
	return a.Class, nil
}

