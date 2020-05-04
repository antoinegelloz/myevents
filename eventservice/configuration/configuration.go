package configuration

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/agelloz/myevents/eventservice/persistence"
)

var (
	// DBTypeDefault is the default DB
	DBTypeDefault = persistence.DBType("mongodb")
	// DBConnectionDefault is the default connection
	DBConnectionDefault = "mongodb://127.0.0.1"
	// EndpointDefault is the default endpoint listening to HTTP
	EndpointDefault = "localhost:8181"
	// TLSEndpointDefault is the default endpoint listening to HTTPS
	TLSEndpointDefault = "localhost:9191"
	// AMPQURLDefault is the default url for the AMPQ broker
	AMPQURLDefault = "amqp://guest:guest@localhost:5672"
)

// ServiceConfig is
type ServiceConfig struct {
	DBType            persistence.DBType `json:"dbType"`
	DBConnection      string             `json:"dbConnection"`
	Endpoint          string             `json:"endpoint"`
	TLSEndpoint       string             `json:"tlsEndpoint"`
	AMQPMessageBroker string             `json:"amqp_message_broker"`
}

// ExtractConfiguration is
func ExtractConfiguration() (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		EndpointDefault,
		TLSEndpointDefault,
		AMPQURLDefault,
	}

	_ = godotenv.Load()
	if os.Getenv("DBTYPE") != "" {
		conf.DBType = persistence.DBType(os.Getenv("DBTYPE"))
	}
	if os.Getenv("DBCONNECTION") != "" {
		conf.DBConnection = os.Getenv("DBCONNECTION")
	}
	if os.Getenv("ENDPOINT") != "" {
		conf.Endpoint = os.Getenv("ENDPOINT")
	}
	if os.Getenv("TLSENDPOINT") != "" {
		conf.TLSEndpoint = os.Getenv("TLSENDPOINT")
	}
	if os.Getenv("AMQP_MESSAGE_BROKER") != "" {
		conf.AMQPMessageBroker = os.Getenv("AMQP_MESSAGE_BROKER")
	}
	log.Printf("%+v\n", conf)
	return conf, nil
}
