// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCategory = "category"

// Category 菜品及套餐分类
type Category struct {
	ID         int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	Type       int32     `gorm:"column:type;type:int;comment:类型   1 菜品分类 2 套餐分类" json:"type"`              // 类型   1 菜品分类 2 套餐分类
	Name       string    `gorm:"column:name;type:varchar(32);not null;comment:分类名称" json:"name"`           // 分类名称
	Sort       int32     `gorm:"column:sort;type:int;not null;comment:顺序" json:"sort"`                     // 顺序
	Status     int32     `gorm:"column:status;type:int;comment:分类状态 0:禁用，1:启用" json:"status"`              // 分类状态 0:禁用，1:启用
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:创建时间" json:"create_time"`         // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;comment:更新时间" json:"update_time"`         // 更新时间
	CreateUser int64     `gorm:"column:create_user;type:bigint;comment:创建人" json:"create_user"`            // 创建人
	UpdateUser int64     `gorm:"column:update_user;type:bigint;comment:修改人" json:"update_user"`            // 修改人
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}
