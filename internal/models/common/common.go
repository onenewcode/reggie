package common

import "encoding/json"

type Result struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r Result) Error() string {
	jsonBytes, _ := json.Marshal(r)

	// 将JSON字节转为字符串并打印
	return string(jsonBytes)
}
