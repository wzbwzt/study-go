package main

import (
	"fmt"
	"time"
)

//worker pool（goroutine池）

//func main() {
//	var jobs = make(chan int, 100)
//	var result = make(chan int, 100)
//	//5个工作
//	for i := 0; i < 5; i++ {
//		jobs <- i
//	}
//	//开启3个goroutine
//	for j := 0; j < 3; j++ {
//		go worker(j, jobs, result)
//	}
//	//输出结果
//	for r := 0; r < 5; r++ {
//		<-result
//	}
//}
//
////示例
//func worker(work int, jobs <-chan int, result chan<- int) {
//	for j := range jobs {
//		fmt.Printf("worker:%d;start:%d\n", work, j)
//		time.Sleep(time.Second)
//		fmt.Printf("worker:%d;end:%d\n", work, j)
//		result <- j * 2
//	}
//}

//Task 定义任务对象
type Task struct {
	f func() error //任务要具体执行的业务 叫f
}

//WorkerPool 定义协程池对象
type WorkerPool struct {
	JobChannel   chan *Task
	EntryChannel chan *Task
	MaxNum       int
}

//NewTask 创建一个任务实例
func NewTask(fun func() error) (task *Task) {
	task = &Task{
		f: fun,
	}
	return
}

//Execute 执行任务实体中的方法
func (t *Task) Execute() {
	t.f() //调用任务中已经绑定的方法
}

//NewPool 创建一个协程池实例
func NewPool(maxNum int) (pool *WorkerPool) {
	pool = &WorkerPool{
		JobChannel:   make(chan *Task),
		EntryChannel: make(chan *Task),
		MaxNum:       maxNum,
	}
	return
}

//Worker 协程池创建一个worker让worker去执行任务
func (w *WorkerPool) Worker(workID int) {
	for task := range w.JobChannel { //永久的从jobchannel中拿取任务去执行
		task.Execute()
		fmt.Println("worker ID:", workID, "has execute a task")
	}
}

func (w *WorkerPool) run() {
	//1.根据maxNum来创建对应数据的worker来工作
	for i := 0; i < w.MaxNum; i++ {
		go w.Worker(i)
	}
	//2.从EntryChannel中取任务，将任务发送给jobchannel
	for tast := range w.EntryChannel {
		w.JobChannel <- tast
	}
}

//testTask 测试使用 模拟具体任务
func testTask() error {
	fmt.Println(time.Now().Second())
	return nil
}

func main() {
	//1.创建一些具体任务
	t := NewTask(testTask)
	//2.创建一个workPool实例
	p := NewPool(6)
	//3.将创建的任务流到entryChannel中
	go func() {
		for { //不断的将任务给写道接口通道中
			p.EntryChannel <- t
		}
	}()
	//4.启动pool去执行
	p.run()
}
