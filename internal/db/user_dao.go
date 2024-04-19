package db

import "reggie/internal/models/model"

type userI interface {
	GetByOpenid(op_id *string) *model.User
	Insert(user *model.User) *model.User
	CountByMap(m map[string]interface{}) int64
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
func (*userDao) CountByMap(m map[string]interface{}) int64 {
	var nums int64
	origin_sql := DBEngine
	// 判断是否含有name，有name不为nil，就进行模糊查询。
	if m["begin"] != nil {
		origin_sql = origin_sql.Where(" create_time >?", m["begin"])
	}
	if m["end"] != nil {
		origin_sql = origin_sql.Where(" create_time <?", m["end"])
	}
	origin_sql.Table(model.TableNameUser).Count(&nums)
	return nums
}
