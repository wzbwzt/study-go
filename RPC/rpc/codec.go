package rpc

import (
	"bytes"
	"encoding/gob"
)

//定义数据格式和编解码  //这边使用gob进行编解码

//定义RPC交互的数据格式
type RPCdata struct {
	FuncName string
	Args []interface{}
}


//编码
func Encode(data RPCdata)([]byte,error){
	var buf bytes.Buffer
	//得到字节数组的编码器
	encoder := gob.NewEncoder(&buf)
	//对数据编码
	err := encoder.Encode(data)
	if err != nil {
		return nil,err
	}
	return buf.Bytes(),nil
}


//解码
func Decode(enData []byte)(RPCdata,error){
	buf:=bytes.NewBuffer(enData)
	decoder := gob.NewDecoder(buf)
	var res RPCdata
	err := decoder.Decode(res)
	if err != nil {
		return res ,err
	}
	return res,nil

}