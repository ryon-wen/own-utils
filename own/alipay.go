package own

import (
	"net/url"

	"github.com/smartwalle/alipay/v3"
)

var client *alipay.Client

type AliPay struct {
	AppID        string `json:"app_id"`
	PrivateKey   string `json:"private_key"`
	AliPublicKey string `json:"ali_public_key"`
	NotifyURL    string `json:"notify_url"`
}

func InitAliPay(pay *AliPay) {
	var err error
	client, err = alipay.New(pay.AppID, pay.PrivateKey, false)
	if err != nil {
		panic(err)
	}
	err = client.LoadAliPayPublicKey(pay.AliPublicKey)
	if err != nil {
		panic(err)
	}
}

func GeneralPayURL(p alipay.TradePagePay) (string, error) {
	//var p = alipay.TradePagePay{}
	//p.NotifyURL = notifyURL
	//p.ReturnURL = "http://xxx"
	//p.Subject = subject
	//p.OutTradeNo = outTradeNo
	//p.TotalAmount = totalAmount
	//p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	var res, err = client.TradePagePay(p)
	if err != nil {
		return "", err
	}

	// generate payURL string
	var payURL = res.String()
	return payURL, nil
}

func DecodePayNotify(values url.Values) (*alipay.Notification, error) {
	return client.DecodeNotification(values)
}
