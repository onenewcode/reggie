package service

import (
	"reggie/internal/db"
	"reggie/internal/models/model"
	"time"
)

func AddShoppingCart(shoppingCart model.ShoppingCart) {
	//判断当前商品是否在购物车中
	shoppingCartList := db.ShopCartDao.List(&shoppingCart)
	if shoppingCartList != &[]model.ShoppingCart{} {
		shoppingCart := (*shoppingCartList)[0]
		shoppingCart.Number += 1
		(*shoppingCartList)[0] = shoppingCart
		db.ShopCartDao.UpdateNumberById(&shoppingCart)
	} else {
		//如果不存在，插入数据，数量就是1
		//判断当前添加到购物车的是菜品还是套餐
		dishId := shoppingCart.DishID
		if dishId != 0 {
			dish := db.DisDao.GetById(dishId)
			shoppingCart.Name = dish.Name
			shoppingCart.Image = dish.Image
			shoppingCart.Amount = dish.Price
		} else {
			setmealId := shoppingCart.SetmealID
			setmeal := db.MealDao.GetByID(setmealId)
			shoppingCart.Name = setmeal.Name
			shoppingCart.Image = setmeal.Image
			shoppingCart.Amount = setmeal.Price
		}
		shoppingCart.Number = 1
		shoppingCart.CreateTime = time.Now()
		db.ShopCartDao.Insert(&shoppingCart)
	}
}
func ShowShoppingCart(id int64) *[]model.ShoppingCart {
	shop_cart := model.ShoppingCart{}
	shop_cart.UserID = id
	return db.ShopCartDao.List(&shop_cart)
}
func CleanShoppingCart(userId int64) {
	db.ShopCartDao.DeleteByUserId(userId)
}
