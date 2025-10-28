package converter

import (
	"fmt"
	"github.com/biryanim/wb_tech_calendar/internal/api/calendar/dto"
	"github.com/biryanim/wb_tech_calendar/internal/model"
	"time"
)

func FromCreateEventReq(req *dto.CreateEventRequest) (*model.Event, error) {
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date: %v", err)
	}
	event := &model.Event{
		UserID: req.UserID,
		Title:  req.Title,
		Date:   date,
	}

	err = event.Validate()
	if err != nil {
		return nil, err
	}

	return event, nil
}

func ToEventResp(event *model.Event) *dto.Event {
	return &dto.Event{
		ID:     event.ID,
		UserID: event.UserID,
		Title:  event.Title,
		Date:   event.Date.String(),
	}
}

func FromUpdateEventReq(req *dto.UpdateEventRequest) (*model.Event, error) {
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date: %v", err)
	}
	event := &model.Event{
		ID:     req.ID,
		UserID: req.UserID,
		Title:  req.Title,
		Date:   date,
	}

	err = event.Validate()
	if err != nil {
		return nil, err
	}

	return event, nil
}

func ToEventsResp(events []*model.Event) []*dto.Event {
	result := make([]*dto.Event, 0, len(events))
	for _, event := range events {
		result = append(result, ToEventResp(event))
	}

	return result
}
