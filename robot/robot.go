package robot

import (
	"github.com/Candy_Wheel_Alert/notify"
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf/browser"
	surf "gopkg.in/headzoo/surf.v1"
)

type Robot struct {
	notifier notify.Notifyer
	bow      *browser.Browser
}

type RichList struct {
	Address string
	Wallet  string
}

func NewRobot(notifyer notify.Notifyer) *Robot {
	r := Robot{}
	r.notifier = notifyer
	r.bow = surf.NewBrowser()
	return &r
}

func (r *Robot) ScrapeBitcoinRichList() ([]RichList, error) {
	err := r.bow.Open("https://bitinfocharts.com/top-100-richest-bitcoin-addresses.html")
	richLists := []RichList{}
	if err != nil {
		return richLists, err
	}
	r.bow.Find(`#tblOne > tbody > tr`).Each(func(arg1 int, arg2 *goquery.Selection) {
		richRist := RichList{}
		arg2.Find("a").Each(func(arg3 int, arg4 *goquery.Selection) {
			if arg3 == 1 {
				richRist.Wallet = arg4.Text()
			} else {
				richRist.Address = arg4.Text()
			}
		})
		richLists = append(richLists, richRist)
	})
	return richLists, nil
}
