// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDish = "dish"

// Dish 菜品
type Dish struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	Name        string    `gorm:"column:name;not null;comment:菜品名称" json:"name"`                // 菜品名称
	CategoryID  int64     `gorm:"column:category_id;not null;comment:菜品分类id" json:"categoryId"` // 菜品分类id
	Price       float64   `gorm:"column:price;comment:菜品价格" json:"price"`                       // 菜品价格
	Image       string    `gorm:"column:image;comment:图片" json:"image"`                         // 图片
	Description string    `gorm:"column:description;comment:描述信息" json:"description"`           // 描述信息
	Status      int32     `gorm:"column:status;default:1;comment:0 停售 1 起售" json:"status"`      // 0 停售 1 起售
	CreateTime  time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`           // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`           // 更新时间
	CreateUser  int64     `gorm:"column:create_user;comment:创建人" json:"create_user"`            // 创建人
	UpdateUser  int64     `gorm:"column:update_user;comment:修改人" json:"update_user"`            // 修改人
}

// TableName Dish's table name
func (*Dish) TableName() string {
	return TableNameDish
}
