package db

import "reggie/internal/models/model"

type addressI interface {
	List(address *model.AddressBook) (*[]model.AddressBook, error)
	// 新增地址
	Save(address *model.AddressBook) error
	// 根据id查询
	GetById(id int64) *model.AddressBook
	// 根据id修改地址
	Update(address *model.AddressBook)
	// 根据 用户id修改 是否默认地址
	UpdateIsDefaultByUserId(address *model.AddressBook)
	// 根据id删除地址
	DeleteById(id int64)
}
type addressDao struct {
}

func (*addressDao) List(address *model.AddressBook) (*[]model.AddressBook, error) {
	var list []model.AddressBook

	origin_sql := DBEngine
	if address.UserID != 0 {
		origin_sql = origin_sql.Where("user_id =? ", address.UserID)
	}
	if address.Phone != "" {
		origin_sql = origin_sql.Where("phone=?", address.Phone)
	}
	origin_sql.Find(&list)
	return &list, nil
}
func (*addressDao) Save(address *model.AddressBook) error {
	if err := DBEngine.Create(address).Error; err != nil {
		return err
	} else {
		return nil
	}
}
func (*addressDao) GetById(id int64) *model.AddressBook {
	var address model.AddressBook
	DBEngine.Where("id=?", id).First(&address)
	return &address
}
func (*addressDao) Update(address *model.AddressBook) {
	DBEngine.Updates(&address)
}
func (*addressDao) UpdateIsDefaultByUserId(address *model.AddressBook) {
	DBEngine.Model(&address).Where("user_id =?", address.UserID).Update("is_default", address.IsDefault)
}
func (*addressDao) DeleteById(id int64) {
	DBEngine.Delete(&model.AddressBook{}, id)
}
