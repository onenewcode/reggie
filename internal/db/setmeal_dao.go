package db

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
	"reggie/internal/models/vo"
)

type mealI interface {
	//GetDishItemBySetmealId(id int64)
	Insert(ml *model.Setmeal) (*model.Setmeal, error)
	PageQuery(page *dto.SetmealPageQueryDTO) (*[]model.Setmeal, int64)
	DeleteByID(id int64)
	GetByID(id int64) *model.Setmeal
	Update(ml *model.Setmeal)
	GtByIdWithDish(id int64) (*vo.SetmealVO, error)
	List(ml *model.Setmeal) *[]model.Setmeal
	GetDishItemBySetmealId(setmealId int64)
}
type mealDao struct {
}

func (m *mealDao) Insert(ml *model.Setmeal) (*model.Setmeal, error) {
	if err := DBEngine.Create(ml).Error; err != nil {
		hlog.Error("数据库插入失败")
		return nil, err
	}
	return ml, nil
}
func (m *mealDao) PageQuery(page *dto.SetmealPageQueryDTO) (*[]model.Setmeal, int64) {
	var (
		dish  []model.Setmeal
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
func (m *mealDao) DeleteByID(id int64) {
	ml := &model.Setmeal{}
	DBEngine.Where("id=?", id).Delete(&ml)
}

/**
 * 根据套餐id查询菜品选项
 */
func (m *mealDao) GetByID(id int64) *model.Setmeal {
	ml := &model.Setmeal{}
	DBEngine.Where("id=?", id).First(&ml)
	return ml
}
func (m *mealDao) Update(ml *model.Setmeal) {
	DBEngine.Create(ml)
}
func (m *mealDao) GtByIdWithDish(id int64) (*vo.SetmealVO, error) {
	v := vo.SetmealVO{}
	err := DBEngine.Table(model.TableNameSetmeal).Preload("SetmealDishes").Where("id=?", 36).First(&v).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}
func (m *mealDao) List(ml *model.Setmeal) *[]model.Setmeal {
	var (
		cat   []model.Setmeal
		count int64
	)
	origin_sql := DBEngine
	// 判断是否含有name，有name不为nil，就进行模糊查询。
	if ml.Name != "" {
		origin_sql = origin_sql.Where("name LIKE ?", "%"+ml.Name+"%")
	}
	if ml.CategoryID != 0 {
		origin_sql = origin_sql.Where("categoryId=?", ml.CategoryID)
	}
	origin_sql.Model(&model.Category{}).Count(&count)
	origin_sql.Find(&cat)
	return &cat
}
func (m *mealDao) GetDishItemBySetmealId(setmealId int64) {

}
