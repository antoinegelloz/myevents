package main

import (
	"github.com/agelloz/myevents/bookingservice/service"
	"log"
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
