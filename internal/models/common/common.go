package common

import (
	"encoding/json"
)

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

type PageResult struct {
	Total   int         `json:"total,omitempty"`   //总记录数
	Records interface{} `json:"records,omitempty"` //当前页数据集合

}
