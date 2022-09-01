package main

import (
	"database/sql"
	data "event-handler/models"
	"log"
	"net/http"
	"time"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	_ "github.com/lib/pq"
)

var eventSubscription = &common.Subscription{
	PubsubName: "pubsub",
	Topic:      "event.create",
	Route:      "/event.create",
}
var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	conn := connectToDB()
	if conn == nil {
		panic("Cant connect to db")
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	s := daprd.NewService(":8080")

	if err := s.AddTopicEventHandler(eventSubscription, app.eventHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error: %v", err)
	}
}

func connectToDB() *sql.DB {
	// get from env
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=events sslmode=disable timezone=UTC connect_timeout=5"

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not ready")
			counts++
		} else {
			log.Println("Connected to DB")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
