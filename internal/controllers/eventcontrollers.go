package controllers

import (
	"apollo-counter/internal/models"
	"apollo-counter/internal/repository"
	"time"
)

type EventController interface {
	CreateEvent(models.NewEventRequest) error
}

type eventController struct {
	eventRepo repository.EventRepository
}

func NewEventController(repo repository.EventRepository) EventController {
	return &eventController{
		eventRepo: repo,
	}
}

func (c *eventController) CreateEvent(req models.NewEventRequest) error {

	parsedTime, err := time.Parse(time.RFC3339, req.DateTime)
	if err != nil {
		return err
	}

	event := models.Event{
		Title:       req.Title,
		Description: req.Description,
		DateTime:    parsedTime,
	}

	// insert into psql db
	err = c.eventRepo.CreateEvent(event)
	if err != nil {
		return err
	}

	// retrieve all events from ps db
	err = c.eventRepo.GetAllEvents()
	if err != nil {
		return err
	}

	// add all events to Redis

	return nil
}
