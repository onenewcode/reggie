package db

import (
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
)

type EmployeeDao struct {
}

func (*EmployeeDao) GetByUserName(username string) *model.Employee {
	var emp model.Employee
	DBEngine.Where("username=?", username).First(&emp)
	return &emp
}
func (*EmployeeDao) Insert(emp *model.Employee) {
	DBEngine.Create(emp)
}

func (*EmployeeDao) PageQuery(page *dto.EmployeePageQueryDTO) *[]model.Employee {
	var users []model.Employee
	origin_sql := DBEngine.Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Order("create_time desc")
	// 判断是否含有name，有name不为nil，就进行模糊查询。
	if page.Name == nil {
		origin_sql.Find(&users)
		return &users
	} else {
		origin_sql.Where("name LIKE ?", "%"+*page.Name+"%").Find(&users)
		return &users
	}
}

func (*EmployeeDao) UpdateStatus(emp *model.Employee) {
	DBEngine.Select("status", "update_time", "update_user").Updates(emp)
}
func (*EmployeeDao) GetById(id int64) *model.Employee {
	var emp model.Employee
	DBEngine.Where("id=?", id).First(&emp)
	return &emp
}
func (*EmployeeDao) Update(emp *model.Employee) {
	DBEngine.Updates(emp)
}
