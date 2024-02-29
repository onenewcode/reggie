package dto

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
