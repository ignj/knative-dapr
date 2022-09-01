package main

import (
	"context"
	"encoding/json"
	data "event-handler/models"
	"log"

	"github.com/dapr/go-sdk/service/common"
)

func (app *Config) eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)

	var event data.Event
	err = json.Unmarshal(e.RawData, &event)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	event.Insert(event)

	return false, nil
}
