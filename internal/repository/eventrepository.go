package repository

import (
	"apollo-counter/internal/models"
	"apollo-counter/internal/utils"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type EventRepository interface {
	GetAllEvents() ([]models.Event, error)
	GetAllCachedEvents() ([]models.Event, error)
	CreateEvent(models.Event) error
	BulkEventCache([]byte, string, time.Duration)
	DeleteEvent(string) error
}

type eventRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewEventRepository(db *sql.DB, rdb *redis.Client) EventRepository {
	return &eventRepository{
		db:  db,
		rdb: rdb,
	}
}

func (r *eventRepository) GetAllEvents() ([]models.Event, error) {

	q := `SELECT id, title, description, datetime, created_at, updated_at FROM public.events`
	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.DateTime,
			&event.CreatedAt,
			&event.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning event: %v", err)
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating events: %v", err)
	}

	return events, nil
}

func (r *eventRepository) GetAllCachedEvents() ([]models.Event, error) {
	ctx := context.Background()
	data, err := r.rdb.Get(ctx, "eventsbulk").Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("no events found in cache")
		}
		return nil, fmt.Errorf("failed to get events from cache: %w", err)
	}

	var events []models.Event
	if err := json.Unmarshal([]byte(data), &events); err != nil {
		return nil, fmt.Errorf("failed to unmarshal events data: %w", err)
	}

	return events, nil
}

// create a new event and add it to PostgreSQL
func (r *eventRepository) CreateEvent(event models.Event) error {

	q := `
		INSERT INTO events (
			title,
			description,
			datetime  
		) VALUES (
			$1, $2, $3
		)
	`
	_, err := r.db.Exec(
		q,
		event.Title,
		event.Description,
		event.DateTime,
	)

	if err != nil {
		pErr, _ := err.(*pq.Error)
		if pErr.Code == "23505" && pErr.Constraint == "events_title_key" {
			return utils.ErrDuplicateTitle
		}
		return err
	}

	return nil
}

// bulk cache events in Redis
func (r *eventRepository) BulkEventCache(data []byte, key string, ttl time.Duration) {
	ctx := context.Background()

	// find out why its being listed as a string instead of JSON
	_ = r.rdb.Set(ctx, key, data, ttl)
}

func (r *eventRepository) DeleteEvent(title string) error {

	var deletedTitle string
	q := `DELETE FROM events WHERE title = $1 RETURNING title`
	err := r.db.QueryRow(q, title).Scan(&deletedTitle)
	if err == sql.ErrNoRows {
		return fmt.Errorf("event with title %s not found", title)
	}
	if err != nil {
		return fmt.Errorf("faild to delete event with error: %s", err)
	}

	return nil
}
