package domain

import (
	"time"
)

// History represents the history domain model
type History struct {
	Base
	ID      int64
	At      time.Time
	Asset   *Asset
	Event   *Event
	Value   float64
}

// Create is a method that creates a history
func (h *History) Create(at time.Time, assetName, eventName string, value float64) error {
	h.Asset = &Asset{}
	if err := h.Asset.GetByName(assetName); err != nil {
		return err
	}
	h.Event = &Event{}
	if err := h.Event.GetByName(eventName); err != nil {
		return err
	}
	sql := `INSERT INTO history (at, asset_id, event_id, value) VALUES ($1, $2, $3, $4) RETURNING id`
	row := h.Repo.QueryRow(sql, at, h.Asset.ID, h.Event.ID, value)
	err := h.Repo.Scan(row, &h.ID)
	if err != nil {
		return err
	}
	return nil
}
