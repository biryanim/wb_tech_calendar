package service

import (
	"context"
	"time"

	"github.com/biryanim/wb_tech_calendar/internal/model"
)

// CalendarService defines the business logic interface for calendar event operations.
type CalendarService interface {
	CreateEvent(ctx context.Context, event *model.Event) (*model.Event, error)
	UpdateEvent(ctx context.Context, event *model.Event) (*model.Event, error)
	DeleteEvent(ctx context.Context, eventID, userID int) error
	GetEventsForDay(ctx context.Context, userID int, date time.Time) ([]*model.Event, error)
	GetEventsForWeek(ctx context.Context, userID int, date time.Time) ([]*model.Event, error)
	GetEventsForMonth(ctx context.Context, userID int, date time.Time) ([]*model.Event, error)
}
