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
