package service

import (
	"context"
	"reggie/internal/dal/model"
	"reggie/internal/dal/query"
	"reggie/internal/db"
)

// 条件查询
func ListAddress(address *model.AddressBook) *[]model.AddressBook {
	list, _ := query.Q.AddressBook.WithContext(context.Background()).List(address)
	return list
}

// 新增地址
func SaveAddress(address *model.AddressBook) {
	query.Q.AddressBook.WithContext(context.Background()).Save(address)
}

// 根据id查询
func GetByIdAddress(id int64) (*model.AddressBook, error) {
	addressBook := query.Q.AddressBook.WithContext(context.Background()).GetById(id)
	return addressBook, nil
}

// 根据id修改地址
func UpdateAddress(address *model.AddressBook) {
	query.Q.AddressBook.WithContext(context.Background()).Save(address)
}

func SetDefaultAddress(address *model.AddressBook) {
	//1、将当前用户的所有地址修改为非默认地址 update address_book set is_default = ? where user_id = ?
	address.IsDefault = false
	query.Q.AddressBook.WithContext(context.Background()).UpdateIsDefaultByUserId(address)

	//2、将当前地址改为默认地址 update address_book set is_default = ? where id = ?
	address.IsDefault = true
	query.Q.AddressBook.WithContext(context.Background()).Updates(address)
}

// 根据id删除地址
func DeleteByIdAddress(id int64) {
	db.AddressDA0.DeleteById(id)
}
