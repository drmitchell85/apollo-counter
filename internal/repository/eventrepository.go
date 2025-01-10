package repository

import (
	"apollo-counter/internal/models"
	"apollo-counter/internal/utils"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type EventRepository interface {
	GetAllEvents() error
	CreateEvent(models.Event) error
	SetEventCache()
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

func (r *eventRepository) GetAllEvents() error {

	q := `SELECT id, title, description, datetime, created_at, updated_at FROM public.events`
	res, err := r.db.Query(q)
	if err != nil {
		return err
	}

	// TODO fix rows

	fmt.Printf("res: %+v", res)
	return nil
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

// cache an event in Redis
func (r *eventRepository) SetEventCache() {

}
