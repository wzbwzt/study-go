//go:build ignore
// +build ignore

package main

import "fmt"

//模拟组装2台电脑

//------------------------
//抽象层：
//有显卡Card 方法：display()
//有内存Memory storage()
//Cpu calcalate()
//------------------------

//------------------------
//实现层：
//有Intel公司  产品：Card Memory Cpu
//Kingston公司 产品：Memory
//NVIDIA公司  产品：Card
//------------------------

//------------------------
//逻辑层：
//组装Intel系列的电脑 并运行
//组装Intel Cpu、Kingston Memory、NVIDIA Card混合的电脑并运行
//------------------------

//------------抽象层------------
type Card interface {
	Display()
}
type Memory interface {
	Storage()
}
type Cpu interface {
	Calcalate()
}

type Computer struct {
	ItemCard   Card
	ItemMemory Memory
	ItemCpu    Cpu
}

//NewComputer  初始化一个computer类对象
func NewComputer(cpu Cpu, card Card, memory Memory) (c *Computer) {
	c = &Computer{ //多态性
		ItemCard:   card,
		ItemCpu:    cpu,
		ItemMemory: memory,
	}
	return
}
func (c *Computer) Work() {
	c.ItemCard.Display()
	c.ItemMemory.Storage()
	c.ItemCpu.Calcalate()
}

//------------实现层------------
type IntelCpu struct {
}

func (i *IntelCpu) Calcalate() {

}

type IntelCard struct {
}

func (i *IntelCard) Display() {

}

type IntelMemory struct {
}

func (i *IntelMemory) Storage() {

}

type KingMemory struct {
}

func (k *KingMemory) Storage() {

}

type NVIDCard struct {
}

func (n *NVIDCard) Display() {

}

//------------逻辑层------------
func main() {
	//组装Intel系列的电脑 并运行
	intelComputer := NewComputer(&IntelCpu{}, &IntelCard{}, &IntelMemory{})
	intelComputer.Work()
	//组装Intel Cpu、Kingston Memory、NVIDIA Card混合的电脑并运行
	mixComputer := NewComputer(&IntelCpu{}, &NVIDCard{}, &KingMemory{})
	mixComputer.Work()

	var a int
	//a = 100
	fmt.Println(&a)

	b := map[string]int{"zs": 12}
	fmt.Println(b)

	//var b map[string]int
	//b=make(map[string]int,0)
	//b["zs"] = 100
	//fmt.Println(b) //会panic
}
