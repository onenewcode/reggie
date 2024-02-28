package db

import (
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
)

type dishI interface {
	PageQuery(page *dto.DishPageQueryDTO) (*[]model.Dish, int64)
	Delete(id *int64) *error
	Update(dish *model.Dish)
	UpdateStatus(cat *model.Dish)
	List(tp *int64) *[]model.Dish
}
type dishDao struct {
}

func (*dishDao) Save(dish *model.Dish) {
	DBEngine.Create(dish)
}
func (*dishDao) PageQuery(page *dto.DishPageQueryDTO) (*[]model.Dish, int64) {
	var (
		cat   []model.Dish
		count int64
	)
	origin_sql := DBEngine
	// 判断是否含有name，有name不为nil，就进行模糊查询。
	if page.Name != nil {
		origin_sql = origin_sql.Where("name LIKE ?", "%"+*page.Name+"%")
	}
	if page.CategoryId != nil {
		origin_sql = origin_sql.Where("categoryId=?", page.CategoryId)
	}
	if page.Status != nil {
		origin_sql = origin_sql.Where("status=?", page.CategoryId)
	}
	origin_sql.Model(&model.Category{}).Count(&count)
	origin_sql.Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Order("create_time desc").Find(&cat)
	return &cat, count
}
func (*dishDao) Delete(id *int64) *error {
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
