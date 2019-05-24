package main

import (
	"log"

	"github.com/Candy_Wheel_Alert/check"
	"github.com/Candy_Wheel_Alert/env"
	"github.com/Candy_Wheel_Alert/notify"
	"github.com/Candy_Wheel_Alert/request"
	"github.com/Candy_Wheel_Alert/robot"
	"github.com/robfig/cron"
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
	request := request.NewRequest()

	cron := cron.New()
	cron.AddFunc(config.Pace, func() {

		richLists, err := robot.ScrapeBitcoinRichList()
		if err != nil {
			shutdown <- err
		}

		check := check.NewCheck(*n, *request)

		err = check.Checktran(richLists)
		if err != nil {
			log.Fatal(err)
		}
	})

	<-shutdown
	log.Fatal(shutdown)

}
