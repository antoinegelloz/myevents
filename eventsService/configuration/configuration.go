package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"reactapp/eventsService/persistence"
)

var (
	// DBTypeDefault is the default DB
	DBTypeDefault = persistence.DBTYPE("mongodb")
	// DBConnectionDefault is the default connection
	DBConnectionDefault = "mongodb://127.0.0.1"
	// EndpointDefault is the default endpoint listening to HTTP
	EndpointDefault = "localhost:8181"
	// TLSEndpointDefault is the default endpoint listening to HTTPS
	TLSEndpointDefault = "localhost:9191"
)

// ServiceConfig is
type ServiceConfig struct {
	DBType       persistence.DBTYPE `json:"databasetype"`
	DBConnection string             `json:"dbconnection"`
	Endpoint     string             `json:"api_endpoint"`
	TLSEndpoint  string             `json:"api_tlsendpoint"`
}

// ExtractConfiguration is
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		EndpointDefault,
		TLSEndpointDefault,
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}
	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
