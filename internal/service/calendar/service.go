package calendar

import (
	"context"
	"fmt"
	"github.com/biryanim/wb_tech_calendar/internal/model"
	"github.com/biryanim/wb_tech_calendar/internal/service"
	"sync"
	"time"
)

var _ service.CalendarService = (*serv)(nil)

type serv struct {
	mu         sync.RWMutex
	events     map[int]*model.Event
	nextID     int
	userEvents map[int][]int
}

func New() *serv {
	return &serv{
		events:     make(map[int]*model.Event),
		nextID:     1,
		userEvents: make(map[int][]int),
	}
}

func (s *serv) CreateEvent(ctx context.Context, event *model.Event) (*model.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	event.ID = s.nextID
	s.nextID++

	s.events[event.ID] = event
	s.userEvents[event.UserID] = append(s.userEvents[event.UserID], event.ID)

	return event, nil
}

func (s *serv) UpdateEvent(ctx context.Context, event *model.Event) (*model.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	curEvent, ok := s.events[event.ID]
	if !ok {
		return nil, model.ErrEventNotFound
	}

	if event.UserID != curEvent.UserID {
		return nil, fmt.Errorf("user id not match: %w", model.ErrEventNotFound)
	}

	curEvent.Date = event.Date
	curEvent.Title = event.Title

	return curEvent, nil
}

func (s *serv) DeleteEvent(ctx context.Context, eventID, userID int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	event, ok := s.events[eventID]
	if !ok {
		return model.ErrEventNotFound
	}

	if event.UserID != userID {
		return fmt.Errorf("user id not match: %w", model.ErrEventNotFound)
	}

	delete(s.events, eventID)

	userEventsIDs := s.userEvents[userID]
	for i, id := range userEventsIDs {
		if id == event.ID {
			userEventsIDs = append(userEventsIDs[:i], userEventsIDs[i+1:]...)
			break
		}
	}

	return nil
}

func (s *serv) GetEventsForDay(ctx context.Context, userID int, date time.Time) ([]*model.Event, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	fmt.Println("Start of day", startOfDay, endOfDay)
	return s.getEventsInRange(userID, startOfDay, endOfDay)
}

func (s *serv) GetEventsForWeek(ctx context.Context, userID int, date time.Time) ([]*model.Event, error) {
	startOfWeek := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfWeek := startOfWeek.Add(7 * 24 * time.Hour)

	return s.getEventsInRange(userID, startOfWeek, endOfWeek)
}

func (s *serv) GetEventsForMonth(ctx context.Context, userID int, date time.Time) ([]*model.Event, error) {
	startOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	return s.getEventsInRange(userID, startOfMonth, endOfMonth)
}

func (s *serv) getEventsInRange(userID int, startDate, endDate time.Time) ([]*model.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	events, ok := s.userEvents[userID]
	if !ok {
		return nil, model.ErrEventNotFound
	}

	var result []*model.Event

	for _, eventID := range events {
		event := s.events[eventID]
		fmt.Println(event.Date.After(startDate), event.Date.Before(endDate), event.Date)
		if !event.Date.Before(startDate) && event.Date.Before(endDate) {
			result = append(result, event)
		}
	}

	fmt.Println("Result", result)
	return result, nil
}
