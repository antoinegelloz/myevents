package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	// DB - database name
	DB = "bookingServiceDB"
	// EVENTS - events collection
	EVENTS = "eventsCollection"
	// BOOKINGS - bookings collection
	BOOKINGS = "bookingsCollection"
)

// DBLayer is the MongoDB persistence layer
type DBLayer struct {
	client *mongo.Client
}

// NewDBLayer is a constructor function to obtain a connection session handler to the desired MongoDB
func NewDBLayer(connection string) (*DBLayer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	return &DBLayer{
		client: client,
	}, err
}
