package main

import (
	"log"

	"github.com/agelloz/reach/eventsService/service"
)

func main() {
	httpErrChan, httpsErrChan := service.ServeAPI()
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httpsErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}
