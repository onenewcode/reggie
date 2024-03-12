package db

import "reggie/internal/models/model"

type meal_dishI interface {
	GetSetmealIdsByDishIds(ids *[]int64) *[]int64
	DeleteByDishId(dish_id int64)
	InsertBatch(ml *[]model.SetmealDish)
	DeleteBySetmealId(id int64)
}
type mealDishDao struct {
}

func (*mealDishDao) GetSetmealIdsByDishIds(ids *[]int64) *[]int64 {
	var nums []int64
	DBEngine.Table(model.TableNameSetmealDish).Select("setmeal_id").Where("id IN (?)").Scan(&nums)
	return &nums
}
func (*mealDishDao) DeleteByDishId(dish_id int64) {
	DBEngine.Table(model.TableNameSetmealDish).Where("dish_id = (?)").Scan(&dish_id)
}
func (*mealDishDao) InsertBatch(ml *[]model.SetmealDish) {
	DBEngine.Select("*").Create(ml)
}
func (*mealDishDao) DeleteBySetmealId(id int64) {
	DBEngine.Where("setmeal_id =?", id).Delete(&model.Setmeal{})
}
