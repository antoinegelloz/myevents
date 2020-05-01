package mongodb

import (
	"context"
	"github.com/agelloz/myevents/bookingservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

// AddEvent adds an event
func (mgoLayer *DBLayer) AddEvent(e models.Event) primitive.ObjectID {
	res, err := mgoLayer.client.Database(DB).Collection(EVENTS).InsertOne(context.TODO(), e)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("AddEvent: inserted a single document: ", res.InsertedID)
	return res.InsertedID.(primitive.ObjectID)
}

// DeleteEvent deletes an event
func (mgoLayer *DBLayer) DeleteEvent(e models.Event) {
	deleteResult, err := mgoLayer.client.Database(DB).Collection(EVENTS).DeleteOne(context.TODO(), e)
	log.Printf("DeleteEvent: deleted %v documents\n", deleteResult.DeletedCount)
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteAllEvents deletes all events
func (mgoLayer *DBLayer) DeleteAllEvents() {
	deleteResult, err := mgoLayer.client.Database(DB).Collection(EVENTS).DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DeleteAllEvents: deleted %v documents\n", deleteResult.DeletedCount)
}

// GetEventByID returns an event
func (mgoLayer *DBLayer) GetEventByID(ID string) (e models.Event) {
	eventID, err := primitive.ObjectIDFromHex(ID)
	if err != nil{
		log.Println("GetEventByID: invalid ObjectID: ", ID)
		return
	}
	result := mgoLayer.client.Database(DB).Collection(EVENTS).FindOne(context.Background(), bson.M{"_id": eventID})
	result.Decode(&e)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("GetEventByID: ", e)
	return
}

// GetEventByName returns an event
func (mgoLayer *DBLayer) GetEventByName(name string) (e models.Event) {
	result := mgoLayer.client.Database(DB).Collection(EVENTS).FindOne(context.Background(), bson.M{"name": name})
	result.Decode(&e)
	log.Println("GetEventByName: ", e)
	return
}

// GetAllEvents returns all events
func (mgoLayer *DBLayer) GetAllEvents() (e []models.Event) {
	cursor, err := mgoLayer.client.Database(DB).Collection(EVENTS).Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	err = cursor.Decode(&e)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// AddBooking adds a booking
func (mgoLayer *DBLayer) AddBooking(b models.Booking) primitive.ObjectID {
	res, err := mgoLayer.client.Database(DB).Collection(BOOKINGS).InsertOne(context.TODO(), b)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("AddBooking: inserted a single document: ", res.InsertedID)
	return res.InsertedID.(primitive.ObjectID)
}

// DeleteBooking deletes a booking
func (mgoLayer *DBLayer) DeleteBooking(b models.Booking) {
	deleteResult, err := mgoLayer.client.Database(DB).Collection(BOOKINGS).DeleteOne(context.TODO(), b)
	log.Printf("DeleteBooking: deleted %v documents\n", deleteResult.DeletedCount)
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllBookings returns all bookings
func (mgoLayer *DBLayer) GetAllBookings() (b []models.Booking) {
	cursor, err := mgoLayer.client.Database(DB).Collection(BOOKINGS).Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	err = cursor.Decode(&b)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// GetBookingByID returns an booking
func (mgoLayer *DBLayer) GetBookingByID(ID string) (b models.Booking) {
	bookingID, err := primitive.ObjectIDFromHex(ID)
	if err != nil{
		log.Println("GetBookingByID: invalid ObjectID: ", ID)
		return
	}
	result := mgoLayer.client.Database(DB).Collection(BOOKINGS).FindOne(context.Background(), bson.M{"_id": bookingID})
	result.Decode(&b)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("GetEventByID: ", b)
	return
}
