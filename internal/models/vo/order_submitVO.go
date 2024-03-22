package vo

import "time"

type OrderSubmitVO struct {
	//订单id
	Id int64 `json:"id,omitempty"`
	//订单号
	OrderNumber string `json:"orderNumber,omitempty"`
	//订单金额
	OrderAmount float64 `json:"orderAmount,omitempty"`
	//下单时间
	OrderTime time.Time `json:"orderTime"`
}
