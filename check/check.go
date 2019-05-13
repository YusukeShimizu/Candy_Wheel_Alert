package check

import (
	"fmt"
	"log"
	"time"
	"encoding/json"
	"strconv"

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

type AddressAsset struct {
	NTx           int
	TotalSent     int64
	FinalBalance  int64
}

func NewCheck(notifyer notify.Notifyer,request request.Request) *Check {
	c := Check{}
	c.notifier = notifyer
	c.request = request
	return &c
}

func (c *Check) Checktran(richLists []robot.RichList) {

	addressAssets := make(map[string]AddressAsset)
	ticker := time.NewTicker(10 * time.Second)
	for {
        select {
		case <-ticker.C:
			for _, richList := range richLists {
				url := "https://blockchain.info/rawaddr/" + richList.Address

				resp,err := c.request.GetMethod(url)
				if err != nil {
					log.Fatal(err)
				}
				var address SingleAddress
				if err := json.Unmarshal(resp, &address); err != nil {
					fmt.Println("JSON Unmarshal error:", err)
					return
				}

				eachasset := AddressAsset {
					NTx: address.NTx,
					TotalSent: address.TotalSent,
					FinalBalance: address.FinalBalance,
				}
				
				// mapにデータがない場合、Line通知しないようにexistsを確認
				_, exists := addressAssets[address.Address]
				if(exists){
					if(addressAssets[address.Address] != eachasset){

						notification := "API取得結果に差分が発生しました。\n";
						notification = notification + "Address:" + address.Address + "\n"
						notification = notification + "Last Info\n";
						notification = notification + "n_tx:" + strconv.Itoa(addressAssets[address.Address].NTx) + "\n"
						notification = notification + "total_sent:" + strconv.FormatInt(addressAssets[address.Address].TotalSent,10) + "\n"
						notification = notification + "final_balance:" + strconv.FormatInt(addressAssets[address.Address].FinalBalance,10) + "\n" + "\n"
						notification = notification + "Latest Info\n";
						notification = notification + "n_tx:" + strconv.Itoa(eachasset.NTx) + "\n"
						notification = notification + "total_sent:" + strconv.FormatInt(eachasset.TotalSent,10) + "\n"
						notification = notification + "final_balance:" + strconv.FormatInt(eachasset.FinalBalance,10) + "\n" + "\n"
						notification = notification + "https://www.blockchain.com/ja/btc/address/"  + address.Address;

						c.notifier.Notify(notification)
					}
				}
				//Mapにアドレスにひもづく情報を格納
				addressAssets[address.Address] = eachasset

			}
        }
    }
}
