package configuration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/agelloz/myevents/bookingservice/persistence"
)

var (
	// DBTypeDefault is the default DB
	DBTypeDefault = persistence.DBType("mongodb")
	// DBConnectionDefault is the default connection
	DBConnectionDefault = "mongodb://127.0.0.1"
	// EndpointDefault is the default endpoint listening to HTTP
	EndpointDefault = "localhost:8282"
	// TLSEndpointDefault is the default endpoint listening to HTTPS
	TLSEndpointDefault = "localhost:9292"
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
	SMTPUsername      string             `json:"smtp_username"`
	SMTPPassword      string             `json:"smtp_password"`
	SMTPHost          string             `json:"smtp_host"`
	SMTPAddr          string             `json:"smtp_addr"`
}

// ExtractConfiguration is
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		EndpointDefault,
		TLSEndpointDefault,
		AMPQURLDefault,
		"",
		"",
		"",
		"",
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Println("configuration file not found. Continuing with default or env values:")
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
		if os.Getenv("SMTP_USERNAME") != "" {
			conf.SMTPUsername = os.Getenv("SMTP_USERNAME")
		}
		if os.Getenv("SMTP_PASSWORD") != "" {
			conf.SMTPPassword = os.Getenv("SMTP_PASSWORD")
		}
		if os.Getenv("SMTP_HOST") != "" {
			conf.SMTPHost = os.Getenv("SMTP_HOST")
		}
		if os.Getenv("SMTP_ADDR") != "" {
			conf.SMTPAddr = os.Getenv("SMTP_ADDR")
		}
		log.Printf("%+v\n", conf)
		return conf, nil
	}
	err = json.NewDecoder(file).Decode(&conf)
	log.Printf("%+v\n", conf)
	return conf, err
}
