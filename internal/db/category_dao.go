package db

import (
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
)

type CategoryDao struct {
}

func (*CategoryDao) Save(category *model.Category) {
	DBEngine.Create(category)
}
func (*CategoryDao) PageQuery(page *dto.CategoryPageQueryDTO) (*[]model.Category, int64) {
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
func (*CategoryDao) Delete(id *int64) *error {
	err := DBEngine.Delete(&model.Category{}, id).Error
	if err != nil {
		return &err
	}
	return nil
}
func (*CategoryDao) Update(category *model.Category) {
	DBEngine.Updates(category)
}
func (*CategoryDao) UpdateStatus(cat *model.Category) {
	DBEngine.Select("status", "update_time", "update_user").Updates(cat)
}
