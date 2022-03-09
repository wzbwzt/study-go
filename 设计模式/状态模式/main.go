// +build ignore

package main

import (
	"errors"
	"fmt"
)

//状态模式：一个对象的内部状态变化时改变其行为，是一种根据状态的行为封装
//经典场景：后台系统的工作流，审批流程等
//前台场景：譬如用户角色升级（游戏）,需要多个状态的判读(积分是否足够，是否已经顶级等)，电商系统商家的入住等
//也可以用于命令行场景

//eg. 比如按照某个软件时，流程
//1.输入序列号
//2.是否同意协议

type IState interface {
	Next(*Machine) error
}

//定义状态机
type Machine struct {
	State IState
}

func NewMachine(state IState) *Machine {
	return &Machine{State: state}
}

func (this *Machine) SetState(state IState) {
	this.State = state
}

func (this *Machine) Next() error {
	return this.State.Next(this)
}

//序列号状态
type AuthNumState struct {
}

func (this *AuthNumState) Next(machine *Machine) error {
	//业务逻辑,判断序列号输入
	fmt.Println("输入序列号:")
	id := ""
	fmt.Scanln(&id)
	if id == "123" {

		//往下执行
		machine.SetState(&LicenseState{})

		return nil
	}

	return fmt.Errorf("序列号错误")

}

//协议状态
type LicenseState struct {
}

var OK = errors.New("success")

func (this *LicenseState) Next(mathine *Machine) error {
	fmt.Println("是否同意安装协议")
	res := ""
	fmt.Scanln(&res)
	if res != "y" {
		return fmt.Errorf("必须统一协议条款才可以安装")
	}
	return OK
}

func main() {
	auth := &AuthNumState{}
	machine := NewMachine(auth)

	for {
		err := machine.Next()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		if err == OK {
			break
		}
	}
}
