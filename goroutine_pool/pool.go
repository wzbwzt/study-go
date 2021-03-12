package main

import (
	"sync"
	"sync/atomic"
	"time"
)

type sig struct{}

type f func() error

// Pool accept the tasks from client,it limits the total
// of goroutines to a given number by recycling goroutines.
type Pool struct {
	// capacity of the pool.
	capacity int32

	// running is the number of the currently running goroutines.
	running int32

	// expiryDuration set the expired time (second) of every worker.
	//在空闲队列中的worker的最新一次运行时间与当前时间之差如果大于这个值则表示已过期，定时清理任务会清理掉这个worker
	expiryDuration time.Duration

	// freeSignal is used to notice pool there are available
	// workers which can be sent to work.
	//freeSignal是一个信号，因为Pool开启的worker数量有上限，因此当全部worker都在执行任务的时候，新进来的请求就需要阻塞等待，那当执行完任务的
	//worker被放回Pool之时，如何通知阻塞的请求绑定一个空闲的worker运行呢？freeSignal就是来做这个事情的
	freeSignal chan sig

	// workers is a slice that store the available workers.
	workers []*Worker

	// release is used to notice the pool to closed itself.
	release chan sig

	// lock for synchronous operation
	lock sync.Mutex

	once sync.Once
}

// putWorker puts a worker back into free pool, recycling the goroutines.
func (p *Pool) putWorker(worker *Worker) {
	// 写入回收时间，亦即该worker的最后运行时间
	worker.recycleTime = time.Now()
	p.lock.Lock()
	p.workers = append(p.workers, worker)
	p.lock.Unlock()
	p.freeSignal <- sig{}
}

func (p *Pool) Cap() int {
	return int(p.capacity)

}

// ReSize change the capacity of this pool
func (p *Pool) ReSize(size int) {
	if size < p.Cap() {
		diff := p.Cap() - size
		for i := 0; i < diff; i++ {
			p.getWorker().stop()
		}
	} else if size == p.Cap() {
		return
	}
	atomic.StoreInt32(&p.capacity, int32(size))
}

//定时清理过期的worker
func (p *Pool) monitorAndClear() {
	go func() {
		for {
			// 周期性循环检查过期worker并清理
			time.Sleep(p.expiryDuration)
			currentTime := time.Now()
			p.lock.Lock()
			idleWorkers := p.workers
			n := 0
			for i, w := range idleWorkers {
				// 计算当前时间减去该worker的最后运行时间之差是否符合过期时长
				if currentTime.Sub(w.recycleTime) <= p.expiryDuration {
					//因为采用了LIFO后进先出队列存放空闲worker，所以该队列默认已经是按照worker的最后运行时间由远及近排序
					break
				}
				n = i
				w.stop()
				//指针指向nil，主动释放堆中资源，不用go的gc来回收，减轻gc的压力
				idleWorkers[i] = nil
				p.running--
			}
			//重新将空闲的worker放入到pool的worker队列中
			if n > 0 {
				n++
				p.workers = idleWorkers[n:]
			}
			p.lock.Unlock()
		}
	}()
}

// Submit submit a task to pool
func (p *Pool) Submit(task f) error {
	if len(p.release) > 0 {
		return ErrPoolClosed
	}
	w := p.getWorker()
	w.sendTask(task)
	return nil
}

// getWorker returns a available worker to run the tasks.
func (p *Pool) getWorker() *Worker {
	var w *Worker
	// 标志，表示当前运行的worker数量是否已达容量上限
	waiting := false
	// 涉及从workers队列取可用worker，需要加锁
	p.lock.Lock()
	workers := p.workers
	n := len(workers) - 1
	// 当前worker队列为空(无空闲worker)
	if n < 0 {
		// 运行worker数目已达到该Pool的容量上限，置等待标志
		if p.running >= p.capacity {
			waiting = true
		} else {
			// 否则，运行数目加1
			p.running++
		}
	} else {
		// 有空闲worker，从队列尾部取出一个使用
		<-p.freeSignal
		w = workers[n]
		//主动回收堆中的资源
		workers[n] = nil
		p.workers = workers[:n]
	}
	// 判断是否有worker可用结束，解锁
	p.lock.Unlock()

	if waiting {
		// 阻塞等待直到有空闲worker
		<-p.freeSignal
		p.lock.Lock()
		workers = p.workers
		l := len(workers) - 1
		w = workers[l]
		workers[l] = nil
		p.workers = workers[:l]
		p.lock.Unlock()
	} else if w == nil {
		// 当前无空闲worker但是Pool还没有满，
		// 则可以直接新开一个worker执行任务
		w = &Worker{
			pool: p,
			task: make(chan f),
		}
		w.run()
	}
	return w
}
