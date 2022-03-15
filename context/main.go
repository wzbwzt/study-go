package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Context 使用原则:

- 不要把Context放在结构体中，要以参数的方式传递。
- 以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。
- 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO。
- Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递。
- Context是线程安全的，可以放心的在多个goroutine中传递。
*/

var wg sync.WaitGroup

//TraceCode 追踪码  用户WithValue
type TraceCode string

//context 用来指定goroutine 的取消退出

func f2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worder 2")
		time.Sleep(time.Millisecond * 200)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}
func f(ctx context.Context) {
	go f2(ctx)
LOOP:
	for {
		fmt.Println("worder 1")
		time.Sleep(time.Millisecond * 200)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	defer wg.Done()
}

//WithTimeout()示例
func connectDB(ctx3 context.Context) {
LOOP:
	for {
		fmt.Println("connecting...")
		time.Sleep(time.Millisecond * 10) // // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx3.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	defer wg.Done()

}

//WithValue()示例
func withvalue(ctxP context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctxP.Value(key).(string) // 在子goroutine中获取trace code  .()类型断言
	if !ok {
		fmt.Println("invalid trace code")
		return
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctxP.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	wg.Done()
}
func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	////Background():Go内置函数要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。
	//wg.Add(1)
	//go f(ctx)
	//time.Sleep(time.Second * 2)
	//cancel() //通知子goroutine结束
	//wg.Wait()
	//---------------------------------------------------------------
	//WithDeadline
	//d := time.Now().Add(time.Second * 1)
	//ctx2, cancel2 := context.WithDeadline(context.Background(), d)
	//尽管ctx2会过期，但在任何情况下调用它的cancel2函数都是很好的实践。
	//如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	//defer cancel2()
	//select {
	//case <-ctx2.Done():
	//	fmt.Println(ctx2.Err())
	//case <-time.After(time.Millisecond * 2000):
	//	fmt.Println("working...")
	//}
	//-----------------------------------------------------------------
	//WithTimeout()//通常用于数据库或者网络连接的超时控制。
	//ctx3, cancel3 := context.WithTimeout(context.Background(), time.Millisecond*50) // 设置一个50毫秒的超时
	//wg.Add(1)
	//go connectDB(ctx3)
	//time.Sleep(time.Second * 2)
	//cancel3()
	//wg.Wait()
	//-----------------------------------------------------------------
	ctxP, cancel := context.WithTimeout(context.Background(), time.Millisecond*50) // 设置一个50毫秒的超时
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctxP = context.WithValue(ctxP, TraceCode("TRACE_CODE"), "123123")
	wg.Add(1)
	go withvalue(ctxP)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}
