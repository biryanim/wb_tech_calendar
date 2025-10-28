package dto

type Event struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
	Title  string `json:"title"`
}

type CreateEventRequest struct {
	UserID int    `json:"user_id" binding:"required"`
	Date   string `json:"date" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type UpdateEventRequest struct {
	ID     int    `json:"id" binding:"required"`
	UserID int    `json:"user_id" binding:"required"`
	Date   string `json:"date"`
	Title  string `json:"title"`
}

type DeleteEventRequest struct {
	ID     int `json:"id" binding:"required"`
	UserID int `json:"user_id" binding:"required"`
}
