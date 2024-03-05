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
type DishVO struct {
	id int64 `json:"id,omitempty"`
	//菜品名称
	name string `json:"name,omitempty"`
	//菜品分类id
	categoryId int64 `json:"category_id,omitempty"`
	//菜品价格
	price float64 `json:"price,omitempty"`
	//图片
	image string `json:"image,omitempty"`
	//描述信息
	description string `json:"description,omitempty"`
	//0 停售 1 起售
	status int64 `json:"status,omitempty"`
	//更新时间
	updateTime time.Time `json:"update_time"`
	//分类名称
	categoryName string `json:"category_name,omitempty"`
	//菜品关联的口味
	flavors []*model.DishFlavor `json:"flavors,omitempty"`
}
