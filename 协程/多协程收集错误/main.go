package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//模拟错误信息
func DoJob(index int) (string, error) {
	if rand.Intn(1000)%2 == 0 {
		return "", fmt.Errorf("index:%d, error", index)
	} else {
		return fmt.Sprintf("success"), nil
	}
}

// func main() {
// 	wg := sync.WaitGroup{}

// 	resChan := make(chan interface{})
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func(index int) {
// 			defer wg.Done()

// 			res, err := DoJob(index)
// 			if err != nil {
// 				resChan <- err
// 			} else {
// 				resChan <- res
// 			}
// 		}(i)
// 	}

// 	go func() {
// 		defer close(resChan)
// 		wg.Wait()
// 	}()

// 	errorCount := 0
// 	for v := range resChan {
// 		if err, ok := v.(error); ok {
// 			fmt.Println(err.Error())
// 			errorCount++
// 		} else {
// 			fmt.Println(v)
// 		}

// 		if errorCount == 2 {
// 			break
// 		}
// 	}

// }

//加协程池来控制协程数量

func main() {
	wg := sync.WaitGroup{}

	pool := make(chan struct{}, 10)

	resChan := make(chan interface{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			wg.Add(1)

			pool <- struct{}{} //会死锁，因为其与resChan同属于一个主协程中，此外resChan是无缓存chan，可以通过将resChan 设置为1000（和循环次数一致）,也可以通过将其放入到协程中执行

			go func(index int) {
				defer wg.Done()

				defer func() {
					<-pool
				}()

				res, err := DoJob(index)
				if err != nil {
					resChan <- err
				} else {
					resChan <- res
				}
			}(i)
		}
	}()

	go func() {
		defer close(resChan)
		wg.Wait()
	}()

	errorCount := 0
	for v := range resChan {
		if err, ok := v.(error); ok {
			fmt.Println(err.Error())
			errorCount++
		} else {
			fmt.Println(v)
		}

		if errorCount == 2 {
			break
		}
	}

}
