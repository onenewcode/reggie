package model

import (
	"time"
)

const TableNameShoppingCart = "shopping_cart"

// ShoppingCart 购物车
type ShoppingCart struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	Name       string    `gorm:"column:name;comment:商品名称" json:"name"`                         // 商品名称
	Image      string    `gorm:"column:image;comment:图片" json:"image"`                         // 图片
	UserID     int64     `gorm:"column:user_id;not null;comment:主键" json:"user_id"`            // 主键
	DishID     int64     `gorm:"column:dish_id;comment:菜品id" json:"dish_id"`                   // 菜品id
	SetmealID  int64     `gorm:"column:setmeal_id;comment:套餐id" json:"setmeal_id"`             // 套餐id
	DishFlavor string    `gorm:"column:dish_flavor;comment:口味" json:"dish_flavor"`             // 口味
	Number     int32     `gorm:"column:number;not null;default:1;comment:数量" json:"number"`    // 数量
	Amount     float64   `gorm:"column:amount;not null;comment:金额" json:"amount"`              // 金额
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`           // 创建时间
}

// TableName ShoppingCart's table name
func (*ShoppingCart) TableName() string {
	return TableNameShoppingCart
}
