package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// DB - database name
	DB = "eventServiceDB"
	// EVENTS - events collection
	EVENTS = "eventsCollection"
)

// DBLayer is the MongoDB persistence layer
type DBLayer struct {
	client *mongo.Client
}

// NewDBLayer is a constructor function to obtain a connection session handler to the desired MongoDB
func NewDBLayer(uri string) *DBLayer {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return &DBLayer{
		client: client,
	}
}
