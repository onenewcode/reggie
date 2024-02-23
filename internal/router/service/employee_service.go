package service

import (
	"reggie/internal/db"
	"reggie/internal/models/constant/status_c"
	"reggie/internal/models/model"
	"time"
)

// 添加成功返回true，添加失败返回flase
func SavEmp(emp *model.Employee) bool {
	//设置账号的状态，默认正常状态 1表示正常 0表示锁定
	emp.Status = status_c.ENABLE

	//设置密码，默认密码123456
	emp.Password = status_c.DEFAULT_PASSWORD

	//设置当前记录的创建时间和修改时间
	emp.CreateTime, emp.UpdateTime = time.Now(), time.Now()

	//设置当前记录创建人id和修改人id
	//emp.CreateUser, emp.UpdateUser = 1, 1 //目前是假数据，之后会继续完善
	// 判断是否用户是否重名
	if db.EmpDao.GetByUserName(emp.Username).Username == emp.Username {
		return false
	}
	db.EmpDao.Insert(emp)
	return true
}
