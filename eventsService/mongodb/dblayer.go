package mongodb

import "gopkg.in/mgo.v2"

const (
	// DB - database name
	DB = "myevents"
	// EVENTS - events
	EVENTS = "events"
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
