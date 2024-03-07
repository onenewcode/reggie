package db

import (
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
)

type catI interface {
	Save(category *model.Category)
	PageQuery(page *dto.CategoryPageQueryDTO) (*[]model.Category, int64)
	Delete(id *int64) *error
	Update(category *model.Category)
	UpdateStatus(cat *model.Category)
	List(tp *int64) *[]model.Category
}
type categoryDao struct {
}

func (*categoryDao) Save(category *model.Category) {
	DBEngine.Create(category)
}
func (*categoryDao) PageQuery(page *dto.CategoryPageQueryDTO) (*[]model.Category, int64) {
	var (
		cat   []model.Category
		count int64
	)
	origin_sql := DBEngine
	// 判断是否含有name，有name不为nil，就进行模糊查询。
	if page.Name != nil {
		origin_sql = origin_sql.Where("name LIKE ?", "%"+*page.Name+"%")
	}
	if page.Type != nil {
		origin_sql = origin_sql.Where("type=?", page.Type)
	}
	origin_sql.Model(&model.Category{}).Count(&count)
	origin_sql.Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Order("create_time desc").Find(&cat)
	return &cat, count
}
func (*categoryDao) Delete(id *int64) *error {
	err := DBEngine.Delete(&model.Category{}, id).Error
	if err != nil {
		return &err
	}
	return nil
}
func (*categoryDao) Update(category *model.Category) {
	DBEngine.Updates(category)
}
func (*categoryDao) UpdateStatus(cat *model.Category) {
	DBEngine.Select("status", "update_time", "update_user").Updates(cat)
}
func (*categoryDao) List(tp *int64) *[]model.Category {
	var cat []model.Category
	DBEngine.Where("type=?", tp).Find(&cat)
	return &cat
}
