package mongodb

import (
	"context"
	"fmt"
	"github.com/agelloz/myevents/eventservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

// AddEvent adds an event
func (mgoLayer *DBLayer) AddEvent(e *models.Event) primitive.ObjectID {
	res, err := mgoLayer.client.Database(DB).Collection(EVENTS).InsertOne(context.TODO(), e)
	if err != nil {
		log.Fatalf("AddEvent: %s\n", err)
	}
	log.Printf("AddEvent: inserted a single document: %+v\n", mgoLayer.GetEventByID(res.InsertedID.(primitive.ObjectID).Hex()))
	return res.InsertedID.(primitive.ObjectID)
}

// DeleteEvent deletes an event
func (mgoLayer *DBLayer) DeleteEvent(e *models.Event) {
	deleteResult, err := mgoLayer.client.Database(DB).Collection(EVENTS).DeleteOne(context.TODO(), bson.M{"_id": e.ID})
	if err != nil {
		log.Fatalf("DeleteEvent: %s\n", err)
	}
	log.Printf("DeleteEvent: deleted %v documents: %s\n", deleteResult.DeletedCount, e.ID)
}

// DeleteAllEvents deletes all events
func (mgoLayer *DBLayer) DeleteAllEvents() {
	deleteResult, err := mgoLayer.client.Database(DB).Collection(EVENTS).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatalf("DeleteAllEvent: %s\n", err)
	}
	log.Printf("DeleteAllEvents: deleted %v documents\n", deleteResult.DeletedCount)
}

// GetEventByID returns an event
func (mgoLayer *DBLayer) GetEventByID(ID string) *models.Event {
	eventID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Println("GetEventByID: invalid ObjectID: ", ID)
		return nil
	}
	var e models.Event
	result := mgoLayer.client.Database(DB).Collection(EVENTS).FindOne(context.Background(), bson.M{"_id": eventID})
	err = result.Decode(&e)
	if err != nil {
		log.Printf("GetEventByID: document not found: %+v: %s\n", eventID, err)
		return nil
	}
	log.Printf("GetEventByID: %+v\n", e)
	return &e
}

// GetEventByName returns an event
func (mgoLayer *DBLayer) GetEventByName(name string) *models.Event {
	var e models.Event
	result := mgoLayer.client.Database(DB).Collection(EVENTS).FindOne(context.Background(), bson.M{"name": name})
	err := result.Decode(&e)
	if err != nil {
		log.Printf("GetEventByName: document not found: %+v: %s\n", name, err)
		return nil
	}
	log.Printf("GetEventByName: %+v\n", e)
	return &e
}

// GetAllEvents returns all events
func (mgoLayer *DBLayer) GetAllEvents() (e []*models.Event) {
	cur, err := mgoLayer.client.Database(DB).Collection(EVENTS).Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatalf("GetAllEvent: %s\n", err)
	}
	for cur.Next(context.TODO()) {
		var elem models.Event
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatalf("GetAllEvent: %s\n", err)
		}
		e = append(e, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatalf("GetAllEvent: %s\n", err)
	}
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", e)
	return
}
