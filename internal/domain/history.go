package domain

import (
	"fmt"
	"time"
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
	fds := []string{"asset_id", "event_id", "at", "value"}
	vals := []string{fmt.Sprintf("%d", h.Asset.ID), fmt.Sprintf("%d", h.Event.ID), at.Format(time.DateTime), fmt.Sprintf("%.2f", value)}
	if h.ID, err = h.Repo.InsertAuto(tx, base, "history", &fds, &vals); err != nil {
		return err
	}
	return nil
}
