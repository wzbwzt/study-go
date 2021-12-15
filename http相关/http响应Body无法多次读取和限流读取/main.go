package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func ReadRsp1(rsp *http.Response) {
	byteBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//重新对响应体赋值，（二手信息）
	rsp.Body = io.NopCloser(bytes.NewBuffer(byteBody))

	fmt.Println("read from 1:", string(byteBody))
}

func ReadRsp2(rsp *http.Response) {
	byteBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//重新对响应体赋值，（二手信息）
	rsp.Body = io.NopCloser(bytes.NewBuffer(byteBody))

	fmt.Println("read from 2:", string(byteBody))
}

//##############################################################################
func main() {
	req, err := http.NewRequest("GET", "https://pc.yiyouliao.com/msn/article.html?recId=46516288387945079d0d15868afdd5fd_s&infoId=II01IXZROEU2V2Q", nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Timeout = time.Second * 5
	rsp, err := http.DefaultClient.Do(req)
	defer rsp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	//限流读取
	limitBody, err := io.ReadAll(io.LimitReader(rsp.Body, 100*1))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(limitBody))

	// ReadRsp1(rsp)
	// ReadRsp2(rsp) //第二次读取不到值

}
