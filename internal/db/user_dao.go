package db

import "reggie/internal/models/model"

type userI interface {
	GetByOpenid(op_id *string) *model.User
	Insert(user *model.User) *model.User
}
type userDao struct {
}

func (*userDao) GetByOpenid(op_id *string) *model.User {
	var user model.User
	result := DBEngine.Where("openid=?", op_id).First(&user)
	// 查询不到，返回nil
	if result.RowsAffected == 0 {
		return nil
	}
	return &user
}
func (*userDao) Insert(user *model.User) *model.User {
	DBEngine.Create(user)
	return user
}
