package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/appengine/log"
)

var DingTalkToken = "https://oapi.dingtalk.com/robot/send?access_token=80573028ce2ca559603366d6ca18c38c160d407c49168ed5dadae86b635163b4"

func SendDingMsg(msg string) {
	keyword := "langmy:" //安全用关键字
	msg = keyword + msg
	//请求地址模板
	webHook := DingTalkToken
	content := `{"msgtype": "text","text": {"content": "` + msg + `"}}`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		// handle error
		log.Errorf("创建dingtaklk 请求 ", errors.WithStack(err))
		fmt.Println(err)
	}
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		// handle error
		log.Errorf("发送dingtakl webhook消息异常", errors.WithStack(err))
	}
}
