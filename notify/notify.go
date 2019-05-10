package notify

import (
	"log"

	"github.com/Candy_Wheel_Alert/env"
	"github.com/line/line-bot-sdk-go/linebot"
)

type Notifyer struct {
	line *linebot.Client
	ID   string
}

func NewNotifyer(config *env.Config) (*Notifyer, error) {
	n := Notifyer{}
	var err error
	n.line, err = linebot.New(config.Secret, config.Token)
	if err != nil {
		return &n, err
	}
	n.ID = config.ID
	return &n, nil
}

func (n *Notifyer) Notify(message string) {
	log.Println(message)
	postMessage := linebot.NewTextMessage(message)
	_, err := n.line.PushMessage(n.ID, postMessage).Do()
	if err != nil {
		log.Fatal(err)
	}
}
