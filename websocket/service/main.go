package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"time"
)



type WsServer struct {
	listener net.Listener
	addr     string
	upgrade  *websocket.Upgrader
}

func NewWsServer() *WsServer {
	ws := new(WsServer)
	ws.addr = "0.0.0.0:8080"
	ws.upgrade = &websocket.Upgrader{
		// 指定升级 websocket 握手完成的超时时间
		HandshakeTimeout :time.Second*5,

		// 写数据操作的缓存池，如果没有设置值，write buffers 将会分配到链接生命周期里。
		//WriteBufferPool BufferPool

		//按顺序指定服务支持的协议，如值存在，则服务会从第一个开始匹配客户端的协议。
		//Subprotocols []string

		// 指定 http 的错误响应函数，如果没有设置 Error 则，会生成 http.Error 的错误响应。
		//Error func(w http.ResponseWriter, r *http.Request, status int, reason error)

		// EnableCompression 指定服务器是否应尝试协商每个邮件压缩（RFC 7692）。
		// 将此值设置为true并不能保证将支持压缩。
		// 目前仅支持“无上下文接管”模式
		//EnableCompression bool

		// 指定 io 操作的缓存大小，如果不指定就会自动分配。
		ReadBufferSize:1024,
		WriteBufferSize:1024,

		// 请求检查函数，用于统一的链接检查，以防止跨站点请求伪造。如果不检查，就设置一个返回值为true的函数。
		// 如果请求Origin标头可以接受，CheckOrigin将返回true。 如果CheckOrigin为nil，则使用安全默认值：
		// 如果Origin请求头存在且原始主机不等于请求主机头，则返回false
		CheckOrigin: func(r *http.Request) bool {
			if r.Method != "GET" {
				fmt.Println("method is not GET")
				return false
			}
			if r.URL.Path != "/ws" {
				fmt.Println("path error")
				return false
			}
			return true
		},
	}
	return ws
}





func (w *WsServer) send10(conn *websocket.Conn) {
	for i := 0; i < 1000; i++ {
		data := fmt.Sprintf("hello websocket test from server %v", time.Now().UnixNano())
		err := conn.WriteMessage(1, []byte(data))
		if err != nil {
			fmt.Println("send msg faild ", err)
			return
		}
		time.Sleep(time.Second * 1)
	}
}

//往conn中发送信息
func (w *WsServer) send(conn *websocket.Conn, stopCh chan int) {
	//w.send10(conn)
	for {
		select {
		case <-stopCh:
			fmt.Println("connect closed")
			return
		case <-time.After(time.Second * 1):
			data := fmt.Sprintf("hello websocket(timeStamp: %v)", time.Now().UnixNano())
			err := conn.WriteMessage(1, []byte(data))
			fmt.Println("sending....")
			if err != nil {
				fmt.Println("send msg faild ", err)
				return
			}
		}
	}
}

//从conn中读取信息
func (w *WsServer) connHandle(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()
	stopCh := make(chan int)
	go w.send(conn, stopCh)
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(5000)))
		_, msg, err := conn.ReadMessage()
		if err != nil {
			close(stopCh)
			// 判断是不是超时
			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() {
					fmt.Printf("ReadMessage timeout remote: %v\n", conn.RemoteAddr())
					return
				}
			}
			// 其他错误，如果是 1001 和 1000 就不打印日志
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				fmt.Printf("ReadMessage other remote:%v error: %v \n", conn.RemoteAddr(), err)
			}
			return
		}
		fmt.Println("receive:", string(msg))
	}
}

//WsServer需要实现了Handler这个接口类
func (w *WsServer) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/ws" {
		httpCode := http.StatusInternalServerError
		statusCodeMsg := http.StatusText(httpCode)
		fmt.Println("500:", statusCodeMsg)
		http.Error(rsp, statusCodeMsg, httpCode)
		return
	}
	conn, err := w.upgrade.Upgrade(rsp, req, nil)
	if err != nil {
		fmt.Println("websocket error:", err)
		return
	}
	fmt.Println("client connect :", conn.RemoteAddr())
	go w.connHandle(conn)

}

func (w *WsServer) Start() (err error) {
	w.listener, err = net.Listen("tcp", w.addr)
	if err != nil {
		fmt.Println("net listen error:", err)
		return
	}
	//WsServer实现了Handler这个接口类所以第二个参数可以直接传w
	err = http.Serve(w.listener, w)
	if err != nil {
		fmt.Println("http serve error:", err)
		return
	}

	return nil
}




func main() {
	webS := NewWsServer()
	webS.Start()
}