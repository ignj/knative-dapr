package data

import (
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		Event: Event{},
	}
}

type Models struct {
	Event Event
}

// TODO: try to autogenerate
type Event struct {
	ID        int       `json:"id"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

func (e *Event) Insert(event Event) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	stmt := `insert into public."Event" ("Data", "CreatedAt") values ($1, $2) returning "Id"`

	err := db.QueryRowContext(ctx, stmt,
		event.Data,
		time.Now(),
	).Scan(&newID)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return newID, nil
}
