package db

import "reggie/internal/models/model"

type EmployeeDao struct {
}

func (*EmployeeDao) GetByUserName(username string) *model.Employee {
	var emp model.Employee
	DBEngine.Where("username=?", username).First(&emp)
	return &emp
}
