package main

import "fmt"

//观察者模式（使用较多）:当一个对象被修改时比如状态，则会通知“观察”他的所有对象
//使用场景：订单系统

//观察者
type Observer interface {
	Update(int64) error
}
type Stock struct {
	ID int64
}

func (this *Stock) Update(id int64) error {
	fmt.Printf("库存%d回滚成功\n", this.ID)
	return nil
}

type Issue struct {
	ID int64
}

func (this *Issue) Update(id int64) error {
	fmt.Printf("工单%d已经取消回滚成功\n", this.ID)
	return nil
}

//##############################################################################

//被观察者
type Subject interface {
	Nodify()
	Add(...Observer)
}
type Order struct {
	Id       int64
	Observer []Observer
}

func (this *Order) Cancel(id int64) {
	fmt.Printf("订单%d取消了\n", id)

	this.Notify(id)
}

func (this *Order) Add(os ...Observer) {
	this.Observer = append(this.Observer, os...)
}

//通知 观察者
func (this *Order) Notify(id int64) error {
	for _, f := range this.Observer {
		err := f.Update(id)
		if err != nil {
			return err
		}
	}
	return nil
}

//##############################################################################

func main() {
	//eg. 被观察者（订单）,观察者（库存，工单）,当订单取消订单时，库存和工单相应触发取消

	//创建观察者
	stock := &Stock{ID: 1}
	issue := &Issue{ID: 2}

	//创建被观察者
	order := Order{}
	order.Add(stock, issue)

	order.Cancel(321)

}
