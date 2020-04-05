package main

import (
	"flag"
	"fmt"
	"log"
	"reactapp/eventsService/configuration"
	"reactapp/eventsService/persistence"
	"reactapp/eventsService/service"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	conf, _ := configuration.ExtractConfiguration(*confPath)
	fmt.Println("Connecting to database...")
	dh, _ := persistence.NewPersistenceLayer(conf.DBType, conf.DBConnection)
	fmt.Println("Serving API...")
	httpErrChan, httpsErrChan := service.ServeAPI(conf.Endpoint, conf.TLSEndpoint, dh)
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httpsErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}
