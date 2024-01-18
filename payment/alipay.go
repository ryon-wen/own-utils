package payment

import (
	"fmt"
	"net/url"

	"github.com/smartwalle/alipay/v3"
)

var client *alipay.Client

func InitAliPay(appId, privateKey, aliPublicKey string) {
	var err error
	client, err = alipay.New(appId, privateKey, false)
	if err != nil {
		panic(err)
	}
	err = client.LoadAliPayPublicKey(aliPublicKey)
	if err != nil {
		panic(err)
	}
}

func GeneralPayment(notifyURL, subject, outTradeNo, totalAmount string) (string, error) {
	var p = alipay.TradePagePay{}
	p.NotifyURL = notifyURL
	//p.ReturnURL = "http://xxx"
	p.Subject = subject
	p.OutTradeNo = outTradeNo
	p.TotalAmount = totalAmount
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	var url, err = client.TradePagePay(p)
	if err != nil {
		return "", err
	}

	// generate payURL string
	var payURL = url.String()
	return payURL, nil
}

func NotifyPayment(values url.Values) (string, uint32, error) {
	var notification, err = client.DecodeNotification(values)
	if err != nil {
		// 错误处理
		fmt.Println(err)
		return "", 0, err
	}
	switch notification.TradeStatus {
	case alipay.TradeStatusWaitBuyerPay: //（交易创建，等待买家付款）
		return notification.OutTradeNo, 0, nil
	case alipay.TradeStatusClosed: //（未付款交易超时关闭，或支付完成后全额退款）
		return notification.OutTradeNo, 1, nil
	case alipay.TradeStatusSuccess: //（交易支付成功）
		return notification.OutTradeNo, 2, nil
	case alipay.TradeStatusFinished: //（交易结束，不可退款）
		return notification.OutTradeNo, 3, nil
	default:
		return "", 0, fmt.Errorf("error request")
	}
}
