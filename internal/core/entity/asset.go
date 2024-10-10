package entity

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
	vals := []string{fmt.Sprintf("%d", a.Class.ID), "'" + name + "'", "'" + description + "'", "'" + time.Now().Format(time.DateTime) + "'"}
	if a.ID, err = a.Repo.InsertAuto(tx, baseName, assetTable, &fds, &vals); err != nil {
		return err
	}
	return nil
}

// GetByName is a method that gets an asset by name
func (a *Asset) GetByName(name string, tx interface{}) error {
	cols, vals, err := a.Repo.Get(tx, baseName, assetTable, "name", name)
	if err != nil {
		return err
	}
	if len(*vals) == 0 {
		return nil
	}
	a.ID, err = strconv.ParseInt(*(*vals)[0][(*cols)["id"]], 10, 64)
	if err != nil {
		return err
	}
	a.Name = *(*vals)[0][(*cols)["name"]]
	a.Description = *(*vals)[0][(*cols)["description"]]
	a.CreatedAt, err = time.Parse(time.DateTime, *(*vals)[0][(*cols)["created_at"]])
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
