package main

//easyfsm :一个用go实现的超容易上手的有限状态机。

//它的特点:
//
//使用简单，快速理解。
//对应状态事件只需全局注册一次，不需要多处注册。
//支持不同业务->相同状态值->自定义不同事件处理器(下面会举🌰)

import (
	"fmt"
	"time"

	"github.com/wuqinqiang/easyfsm"
)

var (
	// 业务
	businessName easyfsm.BusinessName = "order"

	// 对应状态
	initState easyfsm.State = 1 // 初始化
	paidState easyfsm.State = 2 // 已付款
	canceled  easyfsm.State = 3 // 已取消

	//对应事件
	paymentOrderEventName easyfsm.EventName = "paymentOrderEventName"
	cancelOrderEventName  easyfsm.EventName = "cancelOrderEventName"
)

type (
	orderParam struct {
		OrderNo string
	}
)

var (
	_ easyfsm.EventObserver = (*NotifyExample)(nil)
	_ easyfsm.EventHook     = (*HookExample)(nil)
)

type (
	NotifyExample struct {
	}
	HookExample struct {
	}
)

func (h HookExample) Before(opt *easyfsm.Param) {
	println("事件执行前")
}

func (h HookExample) After(opt easyfsm.Param, state easyfsm.State, err error) {
	println("事件执行后")
}

func (o NotifyExample) Receive(opt *easyfsm.Param) {
	println("接收到事件变动,发送消息")
}

func init() {
	// 支付订单事件
	entity := easyfsm.NewEventEntity(paymentOrderEventName,
		func(opt *easyfsm.Param) (easyfsm.State, error) {
			param, ok := opt.Data.(orderParam)
			if !ok {
				panic("param err")
			}
			fmt.Printf("param:%+v\n", param)
			// 处理核心业务
			return paidState, nil
		}, easyfsm.WithHook(HookExample{}), easyfsm.WithObservers(NotifyExample{}))

	// 取消订单事件
	cancelEntity := easyfsm.NewEventEntity(cancelOrderEventName,
		func(opt *easyfsm.Param) (easyfsm.State, error) {
			// 处理核心业务
			param, ok := opt.Data.(orderParam)
			if !ok {
				panic("param err")
			}
			fmt.Printf("param:%+v\n", param)
			return canceled, nil
		}, easyfsm.WithHook(HookExample{}))

	// 注册订单状态机
	easyfsm.RegisterStateMachine(businessName,
		initState,
		entity, cancelEntity)
}

func main() {

	// 正常操作

	// 第一步根据业务，以及当前状态生成fsm
	fsm := easyfsm.NewFSM(businessName, initState)

	// 第二步 调用具体
	currentState, err := fsm.Call(paymentOrderEventName,
		easyfsm.WithData(orderParam{OrderNo: "123"}))

	fmt.Printf("[Success]call paymentOrderEventName err:%v\n", err)
	fmt.Printf("[Success]call paymentOrderEventName state:%v\n", currentState)
	time.Sleep(2 * time.Second)
}
