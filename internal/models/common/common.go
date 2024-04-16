package common

import (
	"encoding/json"
	"fmt"
	"reggie/internal/utils"
	"strings"
	"time"
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
	Total   int64       `json:"total,omitempty"`   //总记录数
	Records interface{} `json:"records,omitempty"` //当前页数据集合
}

// 设置时间格式 2006-01-02 15:04:05
type DateTime time.Time

func (t DateTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
func (t *DateTime) UnmarshalJSON(b []byte) error {
	parse, err := time.Parse("2006-01-02 15:04:05", strings.Trim(utils.Bytes2String(b), "\""))
	if err != nil {
		return err
	}
	*t = (DateTime)(parse)
	return err
}
