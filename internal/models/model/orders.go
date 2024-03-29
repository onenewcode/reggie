package model

import (
	"time"
)

const (
	TableNameOrder = "orders"
	/**
	 * 订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消
	 */

	PENDING_PAYMENT = iota + 1
	TO_BE_CONFIRMED
	CONFIRMED
	DELIVERY_IN_PROGRESS
	COMPLETED
	CANCELLED
)
const (
	UN_PAID = iota
	PAID
	REFUND

	serialVersionUID = 1
)

// Order 订单表
type Order struct {
	ID                    int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                                       // 主键
	Number                string    `gorm:"column:number;comment:订单号" json:"number"`                                                            // 订单号
	Status                int32     `gorm:"column:status;not null;default:1;comment:订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消 7退款" json:"status"`      // 订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消 7退款
	UserID                int64     `gorm:"column:user_id;not null;comment:下单用户" json:"user_id"`                                                // 下单用户
	AddressBookID         int64     `gorm:"column:address_book_id;not null;comment:地址id" json:"address_book_id"`                                // 地址id
	OrderTime             time.Time `gorm:"column:order_time;not null;comment:下单时间" json:"order_time"`                                          // 下单时间
	CheckoutTime          time.Time `gorm:"column:checkout_time;comment:结账时间" json:"checkout_time"`                                             // 结账时间
	PayMethod             int32     `gorm:"column:pay_method;not null;default:1;comment:支付方式 1微信,2支付宝" json:"pay_method"`                       // 支付方式 1微信,2支付宝
	PayStatus             int32     `gorm:"column:pay_status;not null;comment:支付状态 0未支付 1已支付 2退款" json:"pay_status"`                            // 支付状态 0未支付 1已支付 2退款
	Amount                float64   `gorm:"column:amount;not null;comment:实收金额" json:"amount"`                                                  // 实收金额
	Remark                string    `gorm:"column:remark;comment:备注" json:"remark"`                                                             // 备注
	Phone                 string    `gorm:"column:phone;comment:手机号" json:"phone"`                                                              // 手机号
	Address               string    `gorm:"column:address;comment:地址" json:"address"`                                                           // 地址
	UserName              string    `gorm:"column:user_name;comment:用户名称" json:"user_name"`                                                     // 用户名称
	Consignee             string    `gorm:"column:consignee;comment:收货人" json:"consignee"`                                                      // 收货人
	CancelReason          string    `gorm:"column:cancel_reason;comment:订单取消原因" json:"cancel_reason"`                                           // 订单取消原因
	RejectionReason       string    `gorm:"column:rejection_reason;comment:订单拒绝原因" json:"rejection_reason"`                                     // 订单拒绝原因
	CancelTime            time.Time `gorm:"column:cancel_time;comment:订单取消时间" json:"cancel_time"`                                               // 订单取消时间
	EstimatedDeliveryTime time.Time `gorm:"column:estimated_delivery_time;comment:预计送达时间" json:"estimated_delivery_time"`                       // 预计送达时间
	DeliveryStatus        bool      `gorm:"column:delivery_status;not null;default:1;comment:配送状态  1立即送出  0选择具体时间" json:"delivery_status"`      // 配送状态  1立即送出  0选择具体时间
	DeliveryTime          time.Time `gorm:"column:delivery_time;comment:送达时间" json:"delivery_time"`                                             // 送达时间
	PackAmount            int32     `gorm:"column:pack_amount;comment:打包费" json:"pack_amount"`                                                  // 打包费
	TablewareNumber       int32     `gorm:"column:tableware_number;comment:餐具数量" json:"tableware_number"`                                       // 餐具数量
	TablewareStatus       bool      `gorm:"column:tableware_status;not null;default:1;comment:餐具数量状态  1按餐量提供  0选择具体数量" json:"tableware_status"` // 餐具数量状态  1按餐量提供  0选择具体数量
}

// TableName Order's table name
func (*Order) TableName() string {
	return TableNameOrder
}
