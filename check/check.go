package check

import (
	"fmt"
	"time"

	"github.com/Candy_Wheel_Alert/notify"
	"github.com/Candy_Wheel_Alert/robot"
)

type Check struct {
	notifier notify.Notifyer
}

func NewCheck(notifyer notify.Notifyer) *Check {
	c := Check{}
	c.notifier = notifyer
	return &c
}

func (c *Check) Checktran(richLists []robot.RichList) {
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
