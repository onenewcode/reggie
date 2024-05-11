package dto

import (
	"encoding/json"
	"reggie/internal/dal/common"
	"time"
)

type OrdersPageQueryDTO struct {
	Page      int       `json:"page,omitempty"`
	PageSize  int       `json:"page_size,omitempty"`
	Number    string    `json:"number,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Status    int       `json:"status,omitempty"`
	BeginTime time.Time `json:"begin_time"`
	EndTime   time.Time `json:"end_time"`
	UserId    int64     `json:"user_id,omitempty"`
}
type tmpJSON OrdersPageQueryDTO
type tmp struct {
	tmpJSON
	BeginTime common.DateTime `json:"begin_time"`
	EndTime   common.DateTime `json:"end_time"`
}

func (o OrdersPageQueryDTO) MarshalJSON() ([]byte, error) {
	p := &tmp{
		tmpJSON:   (tmpJSON)(o),
		BeginTime: common.DateTime(o.BeginTime),
		EndTime:   common.DateTime(o.EndTime),
	}
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return marshal, err
}

func (o *OrdersPageQueryDTO) UnmarshalJSON(b []byte) error {
	var t tmp
	err := json.Unmarshal(b, &t)
	t.tmpJSON.BeginTime, t.tmpJSON.EndTime = (time.Time)(t.BeginTime), (time.Time)(t.EndTime)
	if err != nil {
		return err
	}
	o = (*OrdersPageQueryDTO)(&t.tmpJSON)
	return nil
}
