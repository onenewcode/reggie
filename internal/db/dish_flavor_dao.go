package db

import "reggie/internal/models/model"

type dishFI interface {
	InsertBatch(flavors *[]model.DishFlavor)
	DeleteByDishId(id int64)
	GetByDishId(id int64) *[]model.DishFlavor
}
type dishFDao struct {
}

func (*dishFDao) InsertBatch(flavors *[]model.DishFlavor) {
	DBEngine.Select("*").Create(flavors)
}
func (*dishFDao) DeleteByDishId(id int64) {
	DBEngine.Table(model.TableNameDishFlavor).Where("id=?", id)
}
func (*dishFDao) GetByDishId(id int64) *[]model.DishFlavor {
	var nums []model.DishFlavor

	DBEngine.Where("dish_id=?", id).Find(&nums)
	return &nums
}
