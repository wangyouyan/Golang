package main

import (
	"fmt"
)

type Phone struct {
	PayMap map[string]Pay
}

func (p *Phone) OpenWeChatPay() {
	wechatpay := &WeChatPay{}
	p.PayMap["wechat_pay"] = wechatpay
}

func (p *Phone) OpenAliPay() {
	alipay := &AliPay{}
	p.PayMap["ali_pay"] = alipay
}

func (p *Phone) PayMoney(name string, money float32) (err error) {
	pay, ok := p.PayMap[name]
	if !ok {
		err = fmt.Errorf("不支持[%s]支付方式", name)
		return
	}

	err = pay.Pay(1023, money)
	return
}