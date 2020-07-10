package rpc

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

var (
	wg sync.WaitGroup
)

func TestSession_WriteRead(t *testing.T) {
	data := "hello wzb"
	addr := "127.0.0.1:8002"
	wg.Add(2)
	//写数据
	go func() {
		defer wg.Done()
		listen, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal("net.listen", err)
		}
		conn, err := listen.Accept()
		if err != nil {
			t.Fatal(err)
		}
		newSession := NewSession(conn)
		err = newSession.Write([]byte(data))
		if err != nil {
			t.Fatal(err)
		}
	}()
	//读数据
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Fatal("net.dail", err)
		}
		newSession := NewSession(conn)
		res, err := newSession.Read()
		if err != nil {
			t.Fatal(err)
		}
		if string(res) != data {
			t.Fatal("test failed")
		}
		fmt.Println(string(res))
	}()

	wg.Wait()

}
