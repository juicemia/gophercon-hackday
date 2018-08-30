package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/juicemia/gophercon-hackday/api"
)

func main() {
	log.Info("booting api server...")

	msgs := make(chan string)

	// TODO: Take this out. This is here to fake consuming messages.
	go func() {
		for {
			msg := <-msgs

			log.Infof("got message: %v", msg)
		}
	}()

	log.Fatal(api.NewServer(":9090", msgs))
}
