package main

import (
	"log"

	"github.com/Candy_Wheel_Alert/check"
	"github.com/Candy_Wheel_Alert/env"
	"github.com/Candy_Wheel_Alert/notify"
	"github.com/Candy_Wheel_Alert/request"
	"github.com/Candy_Wheel_Alert/robot"
)

func main() {

	shutdown := make(chan interface{})

	config, err := env.Process()
	if err != nil {
		log.Fatal(err)
	}

	n, err := notify.NewNotifyer(&config)
	if err != nil {
		log.Fatal(err)
	}

	robot := robot.NewRobot(*n)

	richLists, err := robot.ScrapeBitcoinRichList()
	if err != nil {
		shutdown <- err
	}

	request := request.NewRequest()
	check := check.NewCheck(*n, *request)

	go check.Checktran(richLists)

	<-shutdown
	log.Fatal(shutdown)
}
