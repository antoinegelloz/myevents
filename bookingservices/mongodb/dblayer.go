package mongodb

import "gopkg.in/mgo.v2"

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
	session *mgo.Session
}

// NewDBLayer is a constructor function to obtain a connection session handler to the desired MongoDB
func NewDBLayer(connection string) (*DBLayer, error) {
	s, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}
	return &DBLayer{
		session: s,
	}, err
}
