package db

import (
	"reggie/internal/models/model"
)

type shoppingcartI interface {
	List(shoppingCart *model.ShoppingCart) *[]model.ShoppingCart
	UpdateNumberById(shoppingCart *model.ShoppingCart) error
	Insert(shoppingCart *model.ShoppingCart) *model.ShoppingCart
	DeleteByUserId(userId int64)
}
type shoppingcartDao struct {
}

func (*shoppingcartDao) Insert(shoppingCart *model.ShoppingCart) *model.ShoppingCart {
	DBEngine.Select("*").Create(shoppingCart)
	return shoppingCart
}
func (*shoppingcartDao) List(shoppingCart *model.ShoppingCart) *[]model.ShoppingCart {
	var (
		shop []model.ShoppingCart
		//count int64
	)
	origin_sql := DBEngine
	// 判断是否含有name，有name不为nil，就进行模糊查询。
	if shoppingCart.UserID != 0 {
		origin_sql = origin_sql.Where("userId= ?", shoppingCart.UserID)
	}
	if shoppingCart.SetmealID != 0 {
		origin_sql = origin_sql.Where("setmealId =?", shoppingCart.SetmealID)
	}
	if shoppingCart.DishID != 0 {
		origin_sql = origin_sql.Where("dishId=?", shoppingCart.DishID)
	}
	if shoppingCart.DishFlavor != "" {
		origin_sql = origin_sql.Where("dishFlavor =?", shoppingCart.DishFlavor)
	}
	//origin_sql.Model(&model.Category{}).Count(&count)
	origin_sql.Find(&shop)
	return &shop
}

func (*shoppingcartDao) UpdateNumberById(shoppingCart *model.ShoppingCart) error {
	DBEngine.Updates(shoppingCart)
	return nil
}
func (*shoppingcartDao) DeleteByUserId(userId int64) {
	DBEngine.Where("user_id=?", userId).Delete(&model.ShoppingCart{})
}
