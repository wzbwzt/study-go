package model

type Response struct {
	Code int32       `json:"code"`
	Data interface{} `json:"data"`
}
