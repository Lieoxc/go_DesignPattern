package main

import "fmt"

type PayBehavior interface {
	OrderPay(px *PayCtx)
}
type AliPay struct{}

func (*AliPay) OrderPay(px *PayCtx) {
	fmt.Printf("正在使用Ali pay进行支付,%#v \n", px.payParams)
}

// 三方支付
type ThirdPay struct{}

func (*ThirdPay) OrderPay(px *PayCtx) {
	fmt.Printf("正在使用三方支付进行支付，%#v \n", px.payParams)
}

type PayCtx struct {
	payBehavior PayBehavior
	payParams   map[string]interface{}
}

func (px *PayCtx) SetPayBehavior(p PayBehavior) {
	px.payBehavior = p
}
func (px *PayCtx) Pay() {
	px.payBehavior.OrderPay(px)
}

func NewPayCtx(p PayBehavior) *PayCtx {
	// 支付参数，Mock数据
	params := map[string]interface{}{
		"appId": "234fdfdngj4",
		"mchId": 123456,
	}
	return &PayCtx{
		payBehavior: p,
		payParams:   params,
	}
}

func main() {
	//初始化为 aliPay的支付方式
	px := NewPayCtx(&AliPay{})
	px.Pay()

	//切换为第三方支付
	px.SetPayBehavior(&ThirdPay{})
	px.Pay()
}
