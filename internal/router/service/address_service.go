package service

import (
	"reggie/internal/db"
	"reggie/internal/models/model"
)

// 条件查询
func ListAddress(address *model.AddressBook) *[]model.AddressBook {
	list, _ := db.AddressDA0.List(address)
	return list
}

// 新增地址
func SaveAddress(address *model.AddressBook) {
	db.AddressDA0.Save(address)
}

// 根据id查询
func GetByIdAddress(id int64) (*model.AddressBook, error) {
	addressBook := db.AddressDA0.GetById(id)
	return addressBook, nil
}

// 根据id修改地址
func UpdateAddress(address *model.AddressBook) {
	db.AddressDA0.Update(address)
}

func SetDefaultAddress(address *model.AddressBook) {
	//1、将当前用户的所有地址修改为非默认地址 update address_book set is_default = ? where user_id = ?
	address.IsDefault = false
	db.AddressDA0.UpdateIsDefaultByUserId(address)

	//2、将当前地址改为默认地址 update address_book set is_default = ? where id = ?
	address.IsDefault = true
	db.AddressDA0.Update(address)
}

// 根据id删除地址
func DeleteByIdAddress(id int64) {
	db.AddressDA0.DeleteById(id)
}
