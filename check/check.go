package check

import (
	"fmt"
	"log"
	"time"
	"encoding/json"

	"github.com/Candy_Wheel_Alert/notify"
	"github.com/Candy_Wheel_Alert/robot"
	"github.com/Candy_Wheel_Alert/request"
)

type Check struct {
	notifier notify.Notifyer
	request request.Request
}

type SingleAddress struct {
		Hash160       string `json:"hash160"`
		Address       string `json:"address"`
		NTx           int    `json:"n_tx"`
		TotalReceived int64  `json:"total_received"`
		TotalSent     int64  `json:"total_sent"`
		FinalBalance  int64  `json:"final_balance"`
		Txs           []struct {
			Ver    int `json:"ver"`
			Inputs []struct {
				Sequence int64  `json:"sequence"`
				Witness  string `json:"witness"`
				PrevOut  struct {
					Spent             bool `json:"spent"`
					SpendingOutpoints []struct {
						TxIndex int `json:"tx_index"`
						N       int `json:"n"`
					} `json:"spending_outpoints"`
					TxIndex int    `json:"tx_index"`
					Type    int    `json:"type"`
					Addr    string `json:"addr"`
					Value   int64  `json:"value"`
					N       int    `json:"n"`
					Script  string `json:"script"`
				} `json:"prev_out"`
				Script string `json:"script"`
			} `json:"inputs"`
			Weight      int    `json:"weight"`
			BlockHeight int    `json:"block_height"`
			RelayedBy   string `json:"relayed_by"`
			Out         []struct {
				Spent             bool `json:"spent"`
				SpendingOutpoints []struct {
					TxIndex int `json:"tx_index"`
					N       int `json:"n"`
				} `json:"spending_outpoints,omitempty"`
				TxIndex int    `json:"tx_index"`
				Type    int    `json:"type"`
				Addr    string `json:"addr"`
				Value   int    `json:"value"`
				N       int    `json:"n"`
				Script  string `json:"script"`
			} `json:"out"`
			LockTime   int    `json:"lock_time"`
			Result     int    `json:"result"`
			Size       int    `json:"size"`
			BlockIndex int    `json:"block_index"`
			Time       int    `json:"time"`
			TxIndex    int    `json:"tx_index"`
			VinSz      int    `json:"vin_sz"`
			Hash       string `json:"hash"`
			VoutSz     int    `json:"vout_sz"`
			Rbf        bool   `json:"rbf,omitempty"`
		} `json:"txs"`
}

func NewCheck(notifyer notify.Notifyer,request request.Request) *Check {
	c := Check{}
	c.notifier = notifyer
	c.request = request
	return &c
}

func (c *Check) Checktran(richLists []robot.RichList) {

	resp,err := c.request.GetMethod()
	if err != nil {
		log.Fatal(err)
	}
	var address SingleAddress
	if err := json.Unmarshal(resp, &address); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}
	fmt.Println("-----------------------------------")
	fmt.Println(address.Address)
	fmt.Println(address.NTx)
	fmt.Println(address.FinalBalance)

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
