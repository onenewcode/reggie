package vo

import "reggie/internal/dal/model"

type OrderVO struct {

	//订单菜品信息
	OrderDishes string `json:"orderDishes,omitempty"`
	//订单详情
	OrderDetailList *[]model.OrderDetail `json:"orderDetailList,omitempty"`
}
