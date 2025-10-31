package main

import (
	"log"

	"github.com/biryanim/wb_tech_calendar/internal/api/middleware"

	"github.com/biryanim/wb_tech_calendar/internal/service/calendar"
	"github.com/gin-gonic/gin"

	calendarImpl "github.com/biryanim/wb_tech_calendar/internal/api/calendar"
	"github.com/biryanim/wb_tech_calendar/internal/config"
)

const (
	envFilePath = ".env"
)

func main() {
	err := config.Load(envFilePath)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	httpConfig, err := config.NewHTTPConfig()
	if err != nil {
		log.Fatalf("load http config: %v", err)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware())

	calendarService := calendar.New()
	calendarImpl := calendarImpl.New(calendarService)

	r.POST("/create_event", calendarImpl.CreateEvent)
	r.POST("/update_event", calendarImpl.UpdateEvent)
	r.POST("/delete_event", calendarImpl.DeleteEvent)

	r.GET("/events_for_day", calendarImpl.GetEventsForDay)
	r.GET("/events_for_week", calendarImpl.GetEventsForWeek)
	r.GET("/events_for_month", calendarImpl.GetEventsForMonth)

	if err = r.Run(httpConfig.Address()); err != nil {
		log.Fatal(err)
	}
}
