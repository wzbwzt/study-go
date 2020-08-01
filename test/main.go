package main

import (
	"time"
)

/*
继承的优缺点
优点
简单，直观，关系在编译时静态定义；
被复用的实现易于修改，派生类可以覆盖基类的实现。
缺点
无法在运行时变更从基类继承的实现；
派生类的部分实现通常定义在基类中（派生类可以拓展基类的属性和行为）；
基类的实现细节直接暴露给派生类，破坏了封装；
基类的任何变更都强制子类进行变更（除非子类重写了该方法）。
组合的优缺点
优点
可以在运行时动态对象的行为；
保持类的封装以专注于单一业务；
只通过接口来访问对象，不会破坏封装；
减少对象的依赖关系，更加灵活。
缺点
系统的行为将依赖于不同对象，而不是定义在单个类中，不便于管理；
当需要新的行为时，需要不断定义新的对象来满足需求。
*/
/*

任务类型包括按照文件过期时间清理，按照文件夹容量进行清理及将文件上传至服务器，对于这些任务而言，
具有某些共同点：如都是定时执行的，并且支持启动、停止等操作。

通过Executable接口来将任务的共同点抽象为方法，并且由结构体Task（类）实现该接口定义任务的通用行为。
每类任务也定义为一个具体的结构体（类），并且通过组合Task类来复用Task的代码，使其具有通用行为，
对于各类任务的特有行为而言，如按过期时间清理任务需要遍历文件夹筛选出满足过期条件的文件，按照文件
夹容量清理任务需要先统计文件夹的总容量，当总容量大于警戒容量时再按照修改时间对文件列表进行排序，
从过期时间最久的文件开始删除，直至文件夹容量小于安全容量，则通过实现Executable接口定义的方法Execute来定义各自的行为。
*/

type Executable interface {
	Start()
	Execute()
}

type Task struct {
	executor Executable // 实现hook函数的效果：由子类负责编写业务代码
}

func (t *Task) Start() {
	println("Task.Start()")
	// 复用父类代码
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		//t.Execute()         //Task.Execute()
		t.executor.Execute() // 实现hook函数的效果：由子类负责编写业务代码   //CleanTask.Execute()
	}
}

func (t *Task) Execute() {
	println("Task.Execute()")
}

type CleanTask struct {
	Task
}

func (ct *CleanTask) Execute() {
	println("CleanTask.Execute()")
}

func main () {
	cleanTask := &CleanTask{
		Task{},
	}
	cleanTask.executor = cleanTask // 实现hook函数的效果：由子类负责编写业务代码
	cleanTask.Start()
}



