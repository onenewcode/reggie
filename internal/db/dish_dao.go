package db

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"reggie/internal/constant/status_c"
	"reggie/internal/dal/dto"
	"reggie/internal/dal/model"
)

type dishI interface {
	PageQuery(page *dto.DishPageQueryDTO) (*[]model.Dish, int64)
	Delete(id int64) *error
	Update(dish *model.Dish)
	UpdateStatus(cat *model.Dish)
	List(d *model.Dish) (*[]model.Dish, error)
	GetById(id int64) *model.Dish
	Save(dish *model.Dish) *model.Dish
	GetBySetmealId(id int64) []*model.Dish
}
type dishDao struct {
}

func (*dishDao) Save(dish *model.Dish) *model.Dish {

	if err := DBEngine.Select("*").Create(dish); err != nil {
		hlog.Error(err)
		return nil
	}
	return dish
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

// 根据菜品分类查询菜品
func (*dishDao) List(d *model.Dish) (*[]model.Dish, error) {
	var (
		dish []model.Dish
	)
	origin_sql := DBEngine
	// 判断是否含有name，有name不为nil，就进行模糊查询。
	if d.Name != "" {
		origin_sql = origin_sql.Where("name LIKE ?", "%"+d.Name+"%")
	}
	if d.CategoryID != 0 {
		origin_sql = origin_sql.Where("category_id=?", d.CategoryID)
	}
	if d.Status != status_c.ALL {
		origin_sql = origin_sql.Where("status =?", d.Status)
	}
	if err := origin_sql.Order("create_time desc").Find(&dish).Error; err != nil {
		return nil, err
	}
	return &dish, nil
}
func (*dishDao) GetById(id int64) *model.Dish {
	var dish model.Dish
	DBEngine.Where("id=?", id).First(&dish)
	return &dish
}
func (*dishDao) GetBySetmealId(id int64) []*model.Dish {
	var list []*model.Dish
	DBEngine.Joins(model.TableNameSetmealDish, DBEngine.Where(&model.SetmealDish{SetmealID: id})).Find(&list)
	return list
}
