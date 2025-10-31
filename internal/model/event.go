package model

import (
	"errors"
	"time"
)

// Common errors returned by event operations.
var (
	ErrEventNotFound = errors.New("event not found")
	ErrInvalidUserID = errors.New("invalid user id")
	ErrEmptyTitle    = errors.New("empty title")
	ErrInvalidDate   = errors.New("invalid date")
)

// Event represents a calendar event in the domain model.
type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
}

// Validate checks if the Event has valid field values.
func (e Event) Validate() error {
	if e.UserID <= 0 {
		return ErrInvalidUserID
	}

	if len(e.Title) == 0 {
		return ErrEmptyTitle
	}

	if e.Date.IsZero() {
		return ErrInvalidDate
	}

	return nil
}
