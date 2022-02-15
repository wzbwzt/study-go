package main

import (
	"fmt"
	"net/url"

	"golang.org/x/net/websocket"
)

type Client struct {
	Host string
	Path string
}

func NewWebsocketClient(host, path string) *Client {
	return &Client{
		Host: host,
		Path: path,
	}
}

func (c *Client) SendMessage(body []byte) error {
	u := url.URL{Scheme: "ws", Host: c.Host, Path: c.Path}
	//使用golang.org/x/net/websocket来dial发送消息
	ws, err := websocket.Dial(u.String(), "", "http://"+c.Host+"/")
	defer ws.Close() //关闭连接
	if err != nil {
		fmt.Println("websocket dial err:", err)
		return err
	}

	_, err = ws.Write(body)
	if err != nil {
		fmt.Println("websocket Write err:", err)
		return err
	}
	return nil
}

func main() {
	wc := NewWebsocketClient("127.0.0.1:8080", "ws")
	wc.SendMessage([]byte("hello"))
}
