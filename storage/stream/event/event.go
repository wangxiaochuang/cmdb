package event

import "go.mongodb.org/mongo-driver/mongo"

type Event struct {
	database string
	client   *mongo.Client
}

func NewEvent(client *mongo.Client, db string) (*Event, error) {
	return &Event{client: client, database: db}, nil
}
