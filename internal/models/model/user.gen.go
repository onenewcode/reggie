package model

import (
	"time"
)

const TableNameUser = "user"

// User 用户信息
type User struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	Openid     string    `gorm:"column:openid;comment:微信用户唯一标识" json:"openid"`                 // 微信用户唯一标识
	Name       string    `gorm:"column:name;comment:姓名" json:"name"`                           // 姓名
	Phone      string    `gorm:"column:phone;comment:手机号" json:"phone"`                        // 手机号
	Sex        string    `gorm:"column:sex;comment:性别" json:"sex"`                             // 性别
	IDNumber   string    `gorm:"column:id_number;comment:身份证号" json:"id_number"`               // 身份证号
	Avatar     string    `gorm:"column:avatar;comment:头像" json:"avatar"`                       // 头像
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
