package entity

import (
	"time"
)

const (
	historyTable = "history"
)

// History represents the history domain model
type History struct {
	Base
	ID    int64
	At    time.Time
	Asset *Asset
	Event *Event
	Value float64
}

// Create is a method that creates a history
func (h *History) Create(at time.Time, assetName, eventName string, value float64) error {
	tx, err := h.CheckTx(nil)
	if err != nil {
		return err
	}
	h.Asset = &Asset{}
	if err := h.Asset.GetByName(assetName, tx); err != nil {
		return err
	}
	h.Event = &Event{}
	if err := h.Event.GetByName(eventName, tx); err != nil {
		return err
	}
	vls := map[string]interface{}{
		"asset_id": h.Asset.ID,
		"event_id": h.Event.ID,
		"at":       at.Format(time.DateTime),
		"value":    value,
	}
	if h.ID, err = h.Repo.Insert(tx, baseName, historyTable, &vls); err != nil {
		return err
	}
	return nil
}
