package rpc

import (
	"encoding/binary"
	"io"
	"net"
)

//成熟的RPC框架会有自定义的传输协议，这里的网络传输格式定义；
//前面是固定长度的信息头（header(unit32):指定信息体的长度）；后面是边长信息体(data([]byte))

type session struct{
	conn net.Conn
}


func NewSession(conn net.Conn)*session{
	res:=session{
		conn: conn,
	}
	return &res
}

func(s *session)Write(data []byte)error{
	//4字节+数据长度切片
	buf:=make([]byte,4+len(data))
	//写入头部信息，记录数据长度
	//binary 只认固定长度的类型
	binary.BigEndian.PutUint32(buf[:4],uint32(len(data)))
	//写入数据
	copy(buf[4:],data)
	_, err := s.conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

func(s *session)Read()(data []byte,err error){
	//读取头部信息
	header:=make([]byte,4)
	//按头部长度，读取头部数据
	//io.ReadFull 读取指定长度的数据
	_, err = io.ReadFull(s.conn, header)
	if err != nil {
		return
	}
	//读取数据长度
	dataLen:=binary.BigEndian.Uint32(header)
	data=make([]byte,dataLen)
	_, err = io.ReadFull(s.conn, data)
	if err != nil {
		return
	}
	return data,nil
}

