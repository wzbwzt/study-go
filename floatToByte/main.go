package main

import (
	"encoding/binary"
	"fmt"
	"math"
)

	func main(){

	//float32转[]byte
	var sf float32 = 12.12
	bits := math.Float32bits(sf)
	res:=make([]byte,4)
	binary.LittleEndian.PutUint32(res,bits)
	fmt.Printf("%v\n",res) //[133 235 65 65]

	//[]byte转float32类型
	d:=[]byte{133 ,235, 65, 65}
	u := binary.LittleEndian.Uint32(d)
	frombits := math.Float32frombits(u)
	fmt.Println(frombits) //12.12

	//float64转[]byte
	var f64 =12.324
	u64 := math.Float64bits(f64)
	res64:=make([]byte,8)
	binary.LittleEndian.PutUint64(res64,u64)
	fmt.Printf("%v\n",res64) //[217 206 247 83 227 165 40 64]

	//[]byte转float64类型
	b64:=[]byte{217 ,206 ,247 ,83 ,227 ,165 ,40 ,64}
	ui64 := binary.LittleEndian.Uint64(b64)
	float64frombits := math.Float64frombits(ui64)
	fmt.Println(float64frombits) //12.324

		i := 42             // Signed integer
		f := float64(i)     // Float
		u1 := uint32(f)
		fmt.Printf("%T;%T;%T",u1,f,i)
	}

