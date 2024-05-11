package db

import (
	"reggie/internal/dal/dto"
	"reggie/internal/dal/model"
)

type empI interface {
	GetByUserName(username string) *model.Employee
	Insert(emp *model.Employee)

	PageQuery(page *dto.EmployeePageQueryDTO) (*[]model.Employee, int64)

	UpdateStatus(emp *model.Employee)
	GetById(id int64) *model.Employee
	Update(emp *model.Employee)
}
type employeeDao struct {
}

func (*employeeDao) GetByUserName(username string) *model.Employee {
	var emp model.Employee
	DBEngine.Where("username=?", username).First(&emp)
	return &emp
}
func (*employeeDao) Insert(emp *model.Employee) {
	DBEngine.Create(emp)
}

func (*employeeDao) PageQuery(page *dto.EmployeePageQueryDTO) (*[]model.Employee, int64) {
	var (
		users []model.Employee
		count int64
	)
	origin_sql := DBEngine
	// 判断是否含有name，有name不为nil，就进行模糊查询。
	if page.Name != nil {
		origin_sql = origin_sql.Where("name LIKE ?", "%"+*page.Name+"%").Find(&users)
	}
	origin_sql.Model(&model.Employee{}).Count(&count)
	origin_sql.Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Order("create_time desc").Find(&users)
	return &users, count
}

func (*employeeDao) UpdateStatus(emp *model.Employee) {
	DBEngine.Select("status", "update_time", "update_user").Updates(emp)
}
func (*employeeDao) GetById(id int64) *model.Employee {
	var emp model.Employee
	DBEngine.Where("id=?", id).First(&emp)
	return &emp
}
func (*employeeDao) Update(emp *model.Employee) {
	DBEngine.Updates(emp)
}
