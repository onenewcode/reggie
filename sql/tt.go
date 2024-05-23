package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// 设置日期格式解析
type DateTime time.Time

func (t DateTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
func (t *DateTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	parse, err := time.Parse("2006-01-02 15:04:05", strings.Trim(str, "\""))
	if err != nil {
		return err
	}
	t = (*DateTime)(&parse)
	return err
}

type OrderDTO struct {
	BeginTime time.Time `json:"begin_time"`
	EndTime   time.Time `json:"end_time"`
	UserId    int64     `json:"user_id,omitempty"`
}
type tmpJSON OrderDTO
type tmp struct {
	tmpJSON
	BeginTime DateTime `json:"begin_time"`
	EndTime   DateTime `json:"end_time"`
}

func (o OrderDTO) MarshalJSON() ([]byte, error) {
	p := &tmp{
		tmpJSON:   (tmpJSON)(o),
		BeginTime: DateTime(o.BeginTime),
		EndTime:   DateTime(o.EndTime),
	}
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return marshal, err
}

func (o *OrderDTO) UnmarshalJSON(b []byte) error {
	var t tmp
	err := json.Unmarshal(b, &t)
	t.tmpJSON.BeginTime, t.tmpJSON.EndTime = (time.Time)(t.BeginTime), (time.Time)(t.EndTime)
	if err != nil {
		return err
	}
	o = (*OrderDTO)(&t.tmpJSON)
	return nil
}
func main() {
	my := OrderDTO{
		time.Now(),
		time.Now(),
		123,
	}
	str, _ := my.MarshalJSON()
	fmt.Printf("%s", str)
}
