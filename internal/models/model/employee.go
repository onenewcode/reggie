package model

import (
	"time"
)

const TableNameEmployee = "employee"

// Employee 员工信息
type Employee struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`        // 主键
	Name       string    `gorm:"column:name;not null;comment:姓名" json:"name"`                         // 姓名
	Username   string    `gorm:"column:username;not null;comment:用户名" json:"username"`                // 用户名
	Password   string    `gorm:"column:password;not null;comment:密码" json:"password"`                 // 密码
	Phone      string    `gorm:"column:phone;not null;comment:手机号" json:"phone"`                      // 手机号
	Sex        string    `gorm:"column:sex;not null;comment:性别" json:"sex"`                           // 性别
	IDNumber   string    `gorm:"column:id_number;not null;comment:身份证号" json:"idNumber"`              // 身份证号
	Status     int32     `gorm:"column:status;not null;default:1;comment:状态 0:禁用，1:启用" json:"status"` // 状态 0:禁用，1:启用
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`                  // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                  // 更新时间
	CreateUser int64     `gorm:"column:create_user;comment:创建人" json:"create_user"`                   // 创建人
	UpdateUser int64     `gorm:"column:update_user;comment:修改人" json:"update_user"`                   // 修改人
}

// TableName Employee's table name
func (*Employee) TableName() string {
	return TableNameEmployee
}
