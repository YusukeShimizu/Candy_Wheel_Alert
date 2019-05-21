package check

import (
	"encoding/json"
	"time"
	"strconv"

	"github.com/Candy_Wheel_Alert/notify"
	"github.com/Candy_Wheel_Alert/request"
	"github.com/Candy_Wheel_Alert/robot"
)

type Check struct {
	notifier notify.Notifyer
	request  request.Request
}

type Transaction struct {
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
	Time       int64  `json:"time"`
	TxIndex    int    `json:"tx_index"`
	VinSz      int    `json:"vin_sz"`
	Hash       string `json:"hash"`
	VoutSz     int    `json:"vout_sz"`
}

type SingleAddress struct {
	Hash160       string        `json:"hash160"`
	Address       string        `json:"address"`
	NTx           int           `json:"n_tx"`
	TotalReceived int64         `json:"total_received"`
	TotalSent     int64         `json:"total_sent"`
	FinalBalance  int64         `json:"final_balance"`
	Txs           []Transaction `json:"txs"`
}

type AddressAsset struct {
	NTx          int
	TotalSent    int64
	FinalBalance int64
	Txs          []Transaction
}

func NewCheck(notifyer notify.Notifyer, request request.Request) *Check {
	c := Check{}
	c.notifier = notifyer
	c.request = request
	return &c
}

func (c *Check) Checktran(richLists []robot.RichList) error {

	addressAssets := make(map[string]AddressAsset)
	ticker := time.NewTicker(120 * time.Second)
	for {
		select {
		case <-ticker.C:
			for _, richList := range richLists {
				url := "https://blockchain.info/rawaddr/" + richList.Address

				resp, err := c.request.GetMethod(url)
				if err != nil {
					return err
				}
				var address SingleAddress
				if err := json.Unmarshal(resp, &address); err != nil {
					return err
				}

				eachasset := AddressAsset{
					NTx:          address.NTx,
					TotalSent:    address.TotalSent,
					FinalBalance: address.FinalBalance,
					Txs:          address.Txs,
				}
				for _, tx := range address.Txs {
					// do something
					t := time.Unix(tx.Time, 0)
					diff := time.Now().Sub(t)
					if diff.Hours() <= 1 {
						input_judge := false
						for _, input := range tx.Inputs {
							
							// inputに該当のアドレスがあればTrue
							if input.PrevOut.Addr == address.Address {
								input_judge = true
							}
						}
						if input_judge {
							outvalue := 0
							//var outaddrs []string
							outaddrs := make(map[string]int)
							for _, out := range tx.Out {
								if out.Addr != address.Address {
									outvalue += out.Value
									outaddrs[out.Addr] += out.Value
								}
							}

							var outvalue_btc float64 = float64(outvalue)/100000000
							notification := strconv.FormatFloat(outvalue_btc, 'f', 4, 64) + " BTCの送金\n";
							notification = notification + address.Address + "\n";
							notification = notification + "↓\n";
							for addr, _ := range outaddrs{
								notification = notification + addr + "\n"
							}
							notification = notification + "\n";
							notification = notification + "https://www.blockchain.com/ja/btc/tx/"  + tx.Hash  + "\n";
							c.notifier.Notify(notification)
						}
					}
				}
				//Mapにアドレスにひもづく情報を格納
				addressAssets[address.Address] = eachasset
			}
		}
	}
}
