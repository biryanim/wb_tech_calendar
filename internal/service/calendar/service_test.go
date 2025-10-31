package calendar

import (
	"context"
	"testing"
	"time"

	"github.com/biryanim/wb_tech_calendar/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateEvent(t *testing.T) {
	s := New()
	ctx := context.Background()
	event := &model.Event{
		UserID: 1,
		Title:  "test",
		Date:   time.Now(),
	}

	created, err := s.CreateEvent(ctx, event)
	assert.NoError(t, err)
	assert.Equal(t, 1, created.ID)
	assert.Equal(t, event.Title, s.events[created.ID].Title)
	assert.Equal(t, event.Date, s.events[created.ID].Date)
	assert.Len(t, s.userEvents[event.UserID], 1)
}

func TestUpdateEvent(t *testing.T) {
	s := New()
	ctx := context.Background()
	event := &model.Event{
		UserID: 2,
		Title:  "original",
		Date:   time.Now(),
	}

	created, err := s.CreateEvent(ctx, event)
	assert.NoError(t, err)
	updated := &model.Event{
		ID:     created.ID,
		UserID: 2,
		Title:  "updated",
		Date:   created.Date.Add(1 * time.Hour),
	}

	res, err := s.UpdateEvent(ctx, updated)
	assert.NoError(t, err)
	assert.Equal(t, "updated", res.Title)
	assert.True(t, res.Date.Equal(updated.Date))

	updated.UserID = 99
	_, err = s.UpdateEvent(ctx, updated)
	assert.Error(t, err)

	updated.ID = 333
	updated.UserID = 2
	_, err = s.UpdateEvent(ctx, updated)
	assert.Error(t, err)
}

func TestDeleteEvent(t *testing.T) {
	s := New()
	ctx := context.Background()
	event := &model.Event{
		UserID: 1,
		Title:  "test",
		Date:   time.Now(),
	}
	created, err := s.CreateEvent(ctx, event)
	require.NoError(t, err)

	err = s.DeleteEvent(ctx, created.ID, created.UserID)
	assert.NoError(t, err)
	_, ok := s.events[created.ID]
	assert.False(t, ok)

	created2, err := s.CreateEvent(ctx, event)
	err = s.DeleteEvent(ctx, created2.ID, 99)
	assert.Error(t, err)
}

func TestGetEventsForDay(t *testing.T) {
	s := New()
	ctx := context.Background()
	userID := 9
	date := time.Date(2025, 10, 30, 10, 0, 0, 0, time.UTC)

	e1 := &model.Event{
		UserID: userID,
		Title:  "test",
		Date:   time.Date(2025, 10, 30, 15, 0, 0, 0, time.UTC),
	}

	e2 := &model.Event{
		UserID: userID,
		Title:  "other test",
		Date:   time.Date(2025, 10, 29, 23, 0, 0, 0, time.UTC),
	}
	_, err := s.CreateEvent(ctx, e1)
	require.NoError(t, err)
	_, err = s.CreateEvent(ctx, e2)
	require.NoError(t, err)

	events, err := s.GetEventsForDay(ctx, userID, date)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(events))
	assert.Equal(t, e1.Title, events[0].Title)
}

func TestGetEventsForWeek(t *testing.T) {
	s := New()
	ctx := context.Background()
	userID := 3
	start := time.Date(2025, 10, 30, 0, 0, 0, 0, time.UTC)
	mid := start.Add(4 * 24 * time.Hour)
	late := start.Add(8 * 24 * time.Hour)

	e1 := &model.Event{UserID: userID, Title: "test", Date: mid}
	e2 := &model.Event{UserID: userID, Title: "test2", Date: late}

	_, err := s.CreateEvent(ctx, e1)
	require.NoError(t, err)
	_, err = s.CreateEvent(ctx, e2)
	require.NoError(t, err)

	events, err := s.GetEventsForWeek(ctx, userID, start)
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.Equal(t, e1.Title, events[0].Title)
}

func TestGetEventsForMonth(t *testing.T) {
	s := New()
	ctx := context.Background()
	userID := 3
	start := time.Date(2025, 10, 1, 0, 0, 0, 0, time.UTC)
	mid := time.Date(2025, 10, 15, 12, 0, 0, 0, time.UTC)
	nextMonth := time.Date(2025, 11, 3, 0, 0, 0, 0, time.UTC)

	e1 := &model.Event{UserID: userID, Title: "test", Date: mid}
	e2 := &model.Event{UserID: userID, Title: "test2", Date: nextMonth}

	_, err := s.CreateEvent(ctx, e1)
	require.NoError(t, err)
	_, err = s.CreateEvent(ctx, e2)
	require.NoError(t, err)

	events, err := s.GetEventsForMonth(ctx, userID, start)
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.Equal(t, e1.Title, events[0].Title)
}

func TestGetEvents_NotFound(t *testing.T) {
	s := New()
	ctx := context.Background()
	userID := 3
	date := time.Now()

	_, err := s.GetEventsForDay(ctx, userID, date)
	assert.ErrorIs(t, err, model.ErrEventNotFound)

	_, err = s.GetEventsForWeek(ctx, userID, date)
	assert.ErrorIs(t, err, model.ErrEventNotFound)

	_, err = s.GetEventsForMonth(ctx, userID, date)
	assert.ErrorIs(t, err, model.ErrEventNotFound)
}
