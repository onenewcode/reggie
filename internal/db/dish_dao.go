package db

import (
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
)

type dishI interface {
	PageQuery(page *dto.DishPageQueryDTO) (*[]model.Dish, int64)
	Delete(id int64) *error
	Update(dish *model.Dish)
	UpdateStatus(cat *model.Dish)
	List(tp *int64) *[]model.Dish
	GetById(id int64) *model.Dish
	Save(dish *model.Dish)
}
type dishDao struct {
}

func (*dishDao) Save(dish *model.Dish) {
	DBEngine.Create(dish)
}
func (*dishDao) PageQuery(page *dto.DishPageQueryDTO) (*[]model.Dish, int64) {
	var (
		dish  []model.Dish
		count int64
	)
	origin_sql := DBEngine
	// 判断是否含有name，有name不为nil，就进行模糊查询。
	if page.Name != nil {
		origin_sql = origin_sql.Where("name LIKE ?", "%"+*page.Name+"%")
	}
	if page.CategoryId != nil {
		origin_sql = origin_sql.Where("category_id=?", page.CategoryId)
	}
	if page.Status != nil {
		origin_sql = origin_sql.Where("status=?", page.Status)
	}
	origin_sql.Model(&model.Dish{}).Count(&count)
	origin_sql.Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Order("create_time desc").Find(&dish)
	return &dish, count
}
func (*dishDao) Delete(id int64) *error {
	err := DBEngine.Delete(&model.Dish{}, id).Error
	if err != nil {
		return &err
	}
	return nil
}
func (*dishDao) Update(dish *model.Dish) {
	DBEngine.Updates(dish)
}
func (*dishDao) UpdateStatus(cat *model.Dish) {
	DBEngine.Select("status", "update_time", "update_user").Updates(cat)
}
func (*dishDao) List(tp *int64) *[]model.Dish {
	var dish []model.Dish
	DBEngine.Where("type=?", tp).Find(&dish)
	return &dish
}
func (*dishDao) GetById(id int64) *model.Dish {
	var dish model.Dish
	DBEngine.Where("id=?", id).First(&dish)
	return &dish
}
