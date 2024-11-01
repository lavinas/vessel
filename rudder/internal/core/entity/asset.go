package entity

import (
	"time"

	"github.com/lavinas/vessel/internal/port"
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

// NewAsset creates a new Asset
func NewAsset(repo port.Repository) *Asset {
	return &Asset{
		Base: Base{
			Repo: repo,
		},
	}
}

// Create is a method that creates an asset
func (a *Asset) Create(className, name, description string, tx interface{}) error {
	a.Class = NewClass(a.Repo)
	if err := a.Class.GetByName(className, tx); err != nil {
		return err
	}
	vals := map[string]interface{}{
		"class_id":    a.Class.ID,
		"name":        name,
		"description": description,
		"created_at":  time.Now().Format(time.DateTime),
	}
	var err error
	if a.ID, err = a.Repo.Insert(tx, baseName, assetTable, &vals); err != nil {
		return err
	}
	return nil
}

// GetByID is a method that gets an asset by ID
func (a *Asset) GetByID(id int64, tx interface{}) error {
	where := map[string]interface{}{
		"id": id,
	}
	return a.get(&where, tx)
}

// GetByName is a method that gets an asset by name
func (a *Asset) GetByName(name string, tx interface{}) error {
	where := map[string]interface{}{
		"name": name,
	}
	return a.get(&where, tx)
}

// Loaded is a method that checks if an asset is loaded
func (a *Asset) IsLoaded() bool {
	return a.ID != 0
}

// CheckDuplicity is a method that checks if an asset is duplicated
func (a *Asset) CheckDuplicity(name string, tx interface{}) (bool, error) {
	where := map[string]interface{}{
		"name": name,
	}
	vals, err := a.Repo.Get(tx, baseName, assetTable, &where)
	if err != nil {
		return false, err
	}
	return len(*vals) > 0, nil
}

// CheckClassExistence is a method that checks if the asset class exists
func (a *Asset) CheckClassExistence(className string, tx interface{}) (bool, error) {
	a.Class = NewClass(a.Repo)
	if err := a.Class.GetByName(className, tx); err != nil {
		return false, err
	}
	return a.Class.IsLoaded(), nil
}

// get is a method that gets the asset
func (a *Asset) get(where *map[string]interface{}, tx interface{}) error {
	vals, err := a.Repo.Get(tx, baseName, assetTable, where)
	if err != nil {
		return err
	}
	if len(*vals) == 0 {
		return nil
	}
	if err := a.fill(vals); err != nil {
		return err
	}
	a.Class, err = a.getAssetClass((*vals)[0]["class_id"].(int64), tx)
	if err != nil {
		return err
	}
	return nil
}

// fill is a method that fills the asset with the values
func (a *Asset) fill(vals *[]map[string]interface{}) error {
	a.ID = (*vals)[0]["id"].(int64)
	a.Name = (*vals)[0]["name"].(string)
	a.Description = (*vals)[0]["description"].(string)
	createdAt := (*vals)[0]["created_at"].(string)
	var err error
	a.CreatedAt, err = time.Parse(time.DateTime, createdAt)
	if err != nil {
		return err
	}
	return nil
}

// getAssetClass is a method that gets the asset class
func (a *Asset) getAssetClass(id int64, tx interface{}) (*Class, error) {
	a.Class = NewClass(a.Repo)
	a.Class.ID = id
	if err := a.Class.GetByID(a.Class.ID, tx); err != nil {
		return nil, err
	}
	return a.Class, nil
}
