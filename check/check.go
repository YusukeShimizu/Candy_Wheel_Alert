package check

import (
	"fmt"
	"log"
	"time"

	"github.com/Candy_Wheel_Alert/notify"
	"github.com/Candy_Wheel_Alert/robot"
	"github.com/Candy_Wheel_Alert/request"
)

type Check struct {
	notifier notify.Notifyer
	request request.Request
}

func NewCheck(notifyer notify.Notifyer,request request.Request) *Check {
	c := Check{}
	c.notifier = notifyer
	c.request = request
	return &c
}

func (c *Check) Checktran(richLists []robot.RichList) {

	resp,err := c.request.GetM()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp);

	ticker := time.NewTicker(3 * time.Second)
	for {
        select {
		case <-ticker.C:
			for _, richList := range richLists {
				fmt.Println("Address:", richList.Address)
				fmt.Println(richList.Wallet)
			}
        }
    }
}
