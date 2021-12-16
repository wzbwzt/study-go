package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

//扇入：（fan-in）多个channel 读取数据合并输出到一个总的channel里，然后读取出来做进一步处理（对顺序没有要求）
//直到通道关闭
//是一种收敛模式，主要用来收集处理的结果（后台任务）

func fanIn(chs ...<-chan interface{}) <-chan interface{} {
	var res = make(chan interface{})
	var wg = sync.WaitGroup{}
	for _, ch := range chs {
		wg.Add(1)
		go func(c <-chan interface{}) {
			defer wg.Done()

			for v := range c {
				res <- v
			}
		}(ch)
	}

	go func() {
		defer close(res)
		wg.Wait()
	}()
	return res
}

func getuserInfo() <-chan interface{} {
	var res = make(chan interface{})

	go func() {
		defer close(res)
		time.Sleep(time.Second * 5)
		res <- fmt.Sprintf("<h1>getuserinfo:joel</h1>")
	}()
	return res
}

func getuserScore() <-chan interface{} {
	var res = make(chan interface{})

	go func() {
		defer close(res)
		time.Sleep(time.Second * 2)
		res <- fmt.Sprintf("<h1>getuserScore:100</h1>")
	}()
	return res
}

//##############################################################################
func main() {
	e := gin.New()
	e.GET("/user", func(c *gin.Context) {
		c.Writer.Header().Add("Transfer-Encoding", "chunked")

		c.Writer.WriteHeader(http.StatusOK)
		masterChan := fanIn(getuserInfo(), getuserScore())
		for v := range masterChan {
			fmt.Println(v)
			c.Writer.Write([]byte(v.(string)))
			c.Writer.Flush()
		}
	})

	e.Run(":8080")
}
