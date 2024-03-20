package dto

import "reggie/internal/models/model"

type ShoppingCartDTO struct {
	DishId     int64  `json:"dishId,omitempty"`
	SetmealId  int64  `json:"setmealId,omitempty"`
	DishFlavor string `json:"dishFlavor,omitempty"`
}

func (s *ShoppingCartDTO) ToShoppingCart() model.ShoppingCart {
	var sh model.ShoppingCart
	sh.DishID = s.DishId
	sh.SetmealID = s.SetmealId
	sh.DishFlavor = s.DishFlavor
	return sh
}
