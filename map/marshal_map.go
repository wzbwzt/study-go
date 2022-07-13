//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	{
		s, err := marshalResponse(0, "ok", `{"name": "joel", "city": "shenyang"}`)
		if err != nil {
			fmt.Println("marshal response error:", err)
			return
		}
		fmt.Println(s)
		//output:{"code":0,"msg":"ok","result":"{\"name\": \"tony\", \"city\": \"shenyang\"}"}
		//带有转义符，可读性差
	}

	{
		s, err := marshalResponse1(0, "ok", `{"name": "joel", "city": "shenyang"}`)
		if err != nil {
			fmt.Println("marshal response1 error:", err)
			return
		}
		fmt.Println(s)
		//output:{"code":0,"msg":"ok","result":{"name":"joel","city":"shenyang"}}
		//解决了原始json string,marshal时被添加转义符的问题
	}
}

func marshalResponse(code int, msg string, result interface{}) (string, error) {
	m := map[string]interface{}{
		"code":   0,
		"msg":    "ok",
		"result": result,
	}

	b, err := json.Marshal(&m)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

//改进
//json提供了一种RawMessage类型，本质上就是[]byte，我们将json string转换成RawMessage后
//再传给json.Marshal就可以解决转义符的问题
func marshalResponse1(code int, msg string, result interface{}) (string, error) {
	s, ok := result.(string)
	var m = map[string]interface{}{
		"code": 0,
		"msg":  "ok",
	}

	if ok {

		rawData := json.RawMessage(s)
		m["result"] = rawData
	} else {
		m["result"] = result
	}

	b, err := json.Marshal(&m)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
