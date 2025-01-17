package controllers

import (
	"apollo-counter/internal/models"
	"apollo-counter/internal/repository"
	"encoding/json"
	"time"
)

type EventController interface {
	GetAllEvents() ([]models.Event, error)
	CreateEvent(models.NewEventRequest) error
	DeleteEvent(models.DeleteEventRequest) error
}

type eventController struct {
	eventRepo repository.EventRepository
}

func NewEventController(repo repository.EventRepository) EventController {
	return &eventController{
		eventRepo: repo,
	}
}

func (c *eventController) GetAllEvents() ([]models.Event, error) {

	// return all events from cache
	events, err := c.eventRepo.GetAllCachedEvents()
	if err == nil {
		return events, nil
	}

	// retrieve all events from ps db
	events, err = c.eventRepo.GetAllEvents()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(events)
	if err != nil {
		return nil, err
	}

	// cache events to Redis
	key := "eventsbulk"
	ttl := time.Hour * 24
	c.eventRepo.BulkEventCache(data, key, ttl)

	return events, nil
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
	events, err := c.eventRepo.GetAllEvents()
	if err != nil {
		return err
	}

	data, err := json.Marshal(events)
	if err != nil {
		return err
	}

	// cache events to Redis
	key := "eventsbulk"
	ttl := time.Hour * 24
	c.eventRepo.BulkEventCache(data, key, ttl)

	return nil
}

func (c *eventController) DeleteEvent(req models.DeleteEventRequest) error {

	// delete event from postgresql
	err := c.eventRepo.DeleteEvent(req.Title)
	if err != nil {
		return err
	}

	// update redis cache
	// retrieve all events from ps db
	events, err := c.eventRepo.GetAllEvents()
	if err != nil {
		return err
	}

	data, err := json.Marshal(events)
	if err != nil {
		return err
	}

	// cache events to Redis
	key := "eventsbulk"
	ttl := time.Hour * 24
	c.eventRepo.BulkEventCache(data, key, ttl)

	return nil
}
