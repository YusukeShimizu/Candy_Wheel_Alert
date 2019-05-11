package main

import (
	"log"

	"github.com/Candy_Wheel_Alert/env"
	"github.com/Candy_Wheel_Alert/notify"
)

func main() {

	shutdown := make(chan interface{})

	config, err := env.Process()
	if err != nil {
		log.Fatal(err)
	}

	_, err = notify.NewNotifyer(&config)
	if err != nil {
		log.Fatal(err)
	}

	<-shutdown
	log.Fatal(shutdown)

}
