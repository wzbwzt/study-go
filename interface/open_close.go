package main

import "fmt"

//接口设计的开闭原则(以银行业务为例)

//AbstractBanker 抽象一个AbstractBanker业务员
type AbstractBanker interface {
	DoBusiness()  //抽象的接口 业务接口
}

//实现一个架构层（基于抽象的接口来封装 ，就是在不知道具体有哪些Banker的情况下）
func BankerBusi(a AbstractBanker){
	a.DoBusiness() //多态的现象
}


//SaveBanker 存款的Banker
type SaveBanker struct {

}

//TransBanker 转账的Banker
type TransBanker struct {

}

func (s *SaveBanker)DoBusiness(){
	fmt.Println("存款业务")
}
func (s *TransBanker)DoBusiness(){
	fmt.Println("转账业务")
}



func main(){
	BankerBusi(&SaveBanker{})
	BankerBusi(&TransBanker{})
}