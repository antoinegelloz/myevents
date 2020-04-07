package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/agelloz/reach/eventsService/persistence"
	"os"
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
	DBType       persistence.DBTYPE `json:"dbType"`
	DBConnection string             `json:"dbConnection"`
	Endpoint     string             `json:"endpoint"`
	TLSEndpoint  string             `json:"tlsEndpoint"`
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
		fmt.Println("Configuration file not found. Continuing with default values:")
		fmt.Printf("%+v\n", conf)
		return conf, err
	}
	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
