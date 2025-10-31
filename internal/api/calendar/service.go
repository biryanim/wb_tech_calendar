package calendar

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/biryanim/wb_tech_calendar/internal/api/calendar/dto"
	"github.com/biryanim/wb_tech_calendar/internal/converter"
	"github.com/biryanim/wb_tech_calendar/internal/model"
	"github.com/biryanim/wb_tech_calendar/internal/service"
	"github.com/gin-gonic/gin"
)

// Implementation represents the HTTP handler implementation for calendar event operations.
type Implementation struct {
	calendarService service.CalendarService
}

// New creates a new instance of Implementation with the provided calendar service.
func New(calendarService service.CalendarService) *Implementation {
	return &Implementation{calendarService: calendarService}
}

// ErrorResponse represents the structure of error responses returned by the API.
type ErrorResponse struct {
	Error string `json:"error"`
}

// CreateEvent handles POST requests to create a new calendar event.
func (i *Implementation) CreateEvent(c *gin.Context) {
	var req dto.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request" + err.Error()})
		return
	}

	event, err := converter.FromCreateEventReq(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := i.calendarService.CreateEvent(c.Request.Context(), event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": converter.ToEventResp(res)})
}

// UpdateEvent handles POST requests to update an existing calendar event.
func (i *Implementation) UpdateEvent(c *gin.Context) {
	var req dto.UpdateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request" + err.Error()})
		return
	}

	updateEvent, err := converter.FromUpdateEventReq(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := i.calendarService.UpdateEvent(c.Request.Context(), updateEvent)
	if err != nil {
		if errors.Is(err, model.ErrEventNotFound) {
			c.JSON(http.StatusServiceUnavailable, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": converter.ToEventResp(res)})
}

// DeleteEvent handles POST requests to delete a calendar event.
func (i *Implementation) DeleteEvent(c *gin.Context) {
	var req dto.DeleteEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request" + err.Error()})
		return
	}

	if req.ID <= 0 || req.UserID <= 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}

	err := i.calendarService.DeleteEvent(c.Request.Context(), req.ID, req.UserID)
	if err != nil {
		if errors.Is(err, model.ErrEventNotFound) {
			c.JSON(http.StatusServiceUnavailable, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// GetEventsForDay handles GET requests to retrieve all events for a specific day.
func (i *Implementation) GetEventsForDay(c *gin.Context) {
	userIDstr := c.Query("user_id")
	if len(userIDstr) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid user_id"})
		return
	}
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid user_id" + err.Error()})
	}

	dateStr := c.Query("date")
	if len(dateStr) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid date"})
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid date" + err.Error()})
	}

	events, err := i.calendarService.GetEventsForDay(c.Request.Context(), userID, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, converter.ToEventsResp(events))
}

// GetEventsForWeek handles GET requests to retrieve all events for a specific week.
func (i *Implementation) GetEventsForWeek(c *gin.Context) {
	userIDstr := c.Query("user_id")
	if len(userIDstr) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid user_id"})
		return
	}
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid user_id" + err.Error()})
	}

	dateStr := c.Query("date")
	if len(dateStr) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid date"})
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid date" + err.Error()})
	}

	events, err := i.calendarService.GetEventsForWeek(c.Request.Context(), userID, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, converter.ToEventsResp(events))
}

// GetEventsForMonth handles GET requests to retrieve all events for a specific month.
func (i *Implementation) GetEventsForMonth(c *gin.Context) {
	userIDstr := c.Query("user_id")
	if len(userIDstr) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid user_id"})
		return
	}
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid user_id" + err.Error()})
	}

	dateStr := c.Query("date")
	if len(dateStr) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid date"})
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid date" + err.Error()})
	}

	events, err := i.calendarService.GetEventsForMonth(c.Request.Context(), userID, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, converter.ToEventsResp(events))
}
