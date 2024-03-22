package dto

import "time"

type OrdersSubmitDTO struct {

	//地址簿id
	AddressBookId int64 `json:"addressBookId,omitempty"`
	//付款方式
	PayMethod int32 `json:"payMethod,omitempty"`
	//备注
	Remark string `json:"remark,omitempty"`
	//预计送达时间
	EstimatedDeliveryTime time.Time `json:"estimatedDeliveryTime"`
	//配送状态  1立即送出  0选择具体时间
	DeliveryStatus int32 `json:"deliveryStatus,omitempty"`
	//餐具数量
	TablewareNumber int32 `json:"tablewareNumber,omitempty"`
	//餐具数量状态  1按餐量提供  0选择具体数量
	TablewareStatus int32 `json:"tablewareStatus,omitempty"`
	//打包费
	PackAmount int32 `json:"packAmount,omitempty"`
	//总金额
	Amount float64 `json:"amount,omitempty"`
}
