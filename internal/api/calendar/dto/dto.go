package dto

// Event represents a calendar event in API responses.
type Event struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
	Title  string `json:"title"`
}

// CreateEventRequest represents the payload for creating a new calendar event.
type CreateEventRequest struct {
	UserID int    `json:"user_id" binding:"required"`
	Date   string `json:"date" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

// UpdateEventRequest represents the payload for updating an existing calendar event.
type UpdateEventRequest struct {
	ID     int    `json:"id" binding:"required"`
	UserID int    `json:"user_id" binding:"required"`
	Date   string `json:"date"`
	Title  string `json:"title"`
}

// DeleteEventRequest represents the payload for deleting a calendar event.
type DeleteEventRequest struct {
	ID     int `json:"id" binding:"required"`
	UserID int `json:"user_id" binding:"required"`
}
