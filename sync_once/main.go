package main

/*
- 应用场景：高并发的场景下，有些操作只需要执行一次，例如只加载一次配置文件、只关闭一次通道等；
- sync.Once只有一个Do方法：func (o *Once) Do(f func())
- sync.Once内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全；布尔值用来记录初始化是否完成；这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次
*/

import (
	"sync"
)

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{} //初始化操作只会执行一次
	})
	return instance
}
