package main

import (
	//"code.qschou.com/openapi/go-core/library/log/zlog"
	"fmt"
)

//TODO:工厂模式
type Operater interface {
	Operate(int, int) int
}

type WechatRefund struct {
}

type AliRefund struct {
}

type OperateFactory struct {
}

func NewOperateFactory() *OperateFactory {
	return &OperateFactory{}
}

func (this *OperateFactory) CreateOperate(operatename string) Operater {
	switch operatename {
	case "wechat":
		fmt.Println("wechat")
		return &WechatRefund{}
	case "ali":
		fmt.Println("ali")
		return &AliRefund{}
	default:
		panic("无效运算符号")
		return nil
	}
}

//微信退款
func (this *WechatRefund) Operate(rhs int, lhs int) int {
	return rhs + lhs
}

//支付宝退款
func (this *AliRefund) Operate(rhs int, lhs int) int {
	return rhs * lhs
}

func main() {
	Operator := NewOperateFactory().CreateOperate("wechat")
	fmt.Printf("add result is %d\n", Operator.Operate(1, 2))
}
