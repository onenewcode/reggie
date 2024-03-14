package vo

import (
	"reggie/internal/models/model"
	"time"
)

type EmployeeLoginVO struct {
	Id int64 `json:"id,omitempty"`

	UserName string `json:"user_name,omitempty"`

	Name string `json:"name,omitempty"`

	Token string `json:"token,omitempty"`
}

type SetmealVO struct {
	ID            int64               `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`  // 主键
	CategoryID    int64               `gorm:"column:category_id;not null;comment:菜品分类id" json:"category_id"` // 菜品分类id
	Name          string              `gorm:"column:name;not null;comment:套餐名称" json:"name"`                 // 套餐名称
	Price         float64             `gorm:"column:price;not null;comment:套餐价格" json:"price"`               // 套餐价格
	Status        int32               `gorm:"column:status;default:1;comment:售卖状态 0:停售 1:起售" json:"status"`  // 售卖状态 0:停售 1:起售
	Description   string              `gorm:"column:description;comment:描述信息" json:"description"`            // 描述信息
	Image         string              `gorm:"column:image;comment:图片" json:"image"`                          // 图片
	UpdateTime    time.Time           `gorm:"column:update_time;comment:更新时间" json:"update_time"`            // 更新时间
	SetmealDishes []model.SetmealDish `gorm:"foreignKey:setmeal_id"`
}
