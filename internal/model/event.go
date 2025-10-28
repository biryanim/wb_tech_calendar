package model

import (
	"errors"
	"time"
)

var (
	ErrEventNotFound = errors.New("event not found")
	ErrInvalidUserID = errors.New("invalid user id")
	ErrEmptyTitle    = errors.New("empty title")
	ErrInvalidDate   = errors.New("invalid date")
)

type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
}

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
