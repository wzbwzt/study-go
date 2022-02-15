package main

import (
	"sync/atomic"
	"time"
)

// Worker is the actual executor who runs the tasks,
// it starts a goroutine that accepts tasks and
// performs function calls.
type Worker struct {
	// pool who owns this worker.
	pool *Pool

	// task is a job should be done.
	task chan f

	// recycleTime will be update when putting a worker back into queue.
	recycleTime time.Time
}

// run starts a goroutine to repeat the process
// that performs the function calls.
func (w *Worker) run() {
	//atomic.AddInt32(&w.pool.running, 1)
	go func() {
		//监听任务列表，一旦有任务立马取出运行
		for f := range w.task {
			if f == nil {
				atomic.AddInt32(&w.pool.running, -1)
				return
			}
			f()

			//回收复用
			w.pool.putWorker(w)
		}
	}()
}

// stop this worker.
func (w *Worker) stop() {
	w.sendTask(nil)
}

// sendTask sends a task to this worker.
func (w *Worker) sendTask(task f) {
	w.task <- task
}
