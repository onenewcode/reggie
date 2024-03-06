package dto

import (
	"encoding/json"
	"reggie/internal/models/model"
	"time"
)

type EmployeePageQueryDTO struct {
	//员工姓名
	Name *string `json:"name,omitempty" form:"name,omitempty"`
	//页码
	Page int `json:"page,omitempty" form:"page,omitempty"`
	//每页显示记录数
	PageSize int `json:"pageSize,omitempty" form:"pageSize,omitempty"`
}
type CategoryPageQueryDTO struct {
	//员工姓名
	Name *string `json:"name,omitempty" form:"name,omitempty"`
	//页码
	Page int `json:"page,omitempty" form:"page,omitempty"`
	//每页显示记录数
	PageSize int `json:"pageSize,omitempty" form:"pageSize,omitempty"`
	//分类类型 1菜品分类  2套餐分类
	Type *int `json:"type,omitempty" form:"type,omitempty"`
}
type DishDTO struct {
	ID          int64               `json:"id,omitempty"`
	Name        string              `json:"name,omitempty"`
	CategoryID  int64               `json:"categoryId,omitempty"`
	Price       float64             `json:"price,omitempty"`
	Image       string              `json:"image,omitempty"`
	Description string              `json:"description,omitempty"`
	Status      int32               `json:"status,omitempty"`
	flavors     []*model.DishFlavor `json:"flavors,omitempty"`
}

// 如果传入的id不等于nil，
func (d *DishDTO) ToNewDish(id *int64) *model.Dish {
	v, _ := json.Marshal(d)
	var dish model.Dish
	json.Unmarshal(v, &dish)
	dish.CreateUser, dish.UpdateUser = *id, *id
	dish.CreateTime, dish.UpdateTime = time.Now(), time.Now()
	return &dish
}

/*
添加分页id
*/
type DishPageQueryDTO struct {
	Page int `json:"page,omitempty" form:"page,omitempty"`

	PageSize int `json:"pageSize,omitempty" form:"pageSize,omitempty"`

	Name *string `json:"name,omitempty" form:"name,omitempty"`

	//分类id
	CategoryId *int `json:"category_id,omitempty" form:"categoryId,omitempty"`

	//状态 0表示禁用 1表示启用
	Status *int `json:"status,omitempty" form:"status,omitempty"`
}

/**
 * C端用户登录
 */
type UserLoginDTO struct {
	Code string `json:"code,omitempty"`
}

// userDto转换成user，更新用户时间
func (u *UserLoginDTO) ToNewUser() *model.User {
	var us model.User
	us.Openid = u.Code
	us.CreateTime = time.Now()
	return &us
}
