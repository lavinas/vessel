package entity

import (
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
	vals := map[string]interface{}{
		"class_id":   a.Class.ID,
		"name":       name,
		"description": description,
		"created_at": time.Now().Format(time.DateTime),
	} 
	if a.ID, err = a.Repo.Insert(tx, baseName, assetTable, &vals); err != nil {
		return err
	}
	return nil
}

// GetByName is a method that gets an asset by name
func (a *Asset) GetByName(name string, tx interface{}) error {
	where := map[string]interface{}{
		"name": name,
	}
	vals, err := a.Repo.Get(tx, baseName, assetTable, &where)
	if err != nil {
		return err
	}
	if len(*vals) == 0 {
		return nil
	}
	a.ID = (*vals)[0]["id"].(int64)
	a.Name = (*vals)[0]["name"].(string)
	a.Description = (*vals)[0]["description"].(string)
	a.CreatedAt = (*vals)[0]["created_at"].(time.Time)
	a.Class, err = a.getAssetClass((*vals)[0]["class_id"].(int64), tx)
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
