package service

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"reggie/internal/db"
	"reggie/internal/models/common"
	"reggie/internal/models/constant/message_c"
	"reggie/internal/models/constant/status_c"
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
	"reggie/internal/models/vo"
)

// 第一个参数菜品，第二个参数菜品口味数组
func SaveSetMealWithDish(dish *model.Setmeal, flavors *[]model.SetmealDish) error {
	// 	首先插入套餐表
	dh, err := db.MealDao.Insert(dish)
	if err != nil {
		hlog.Error(err)
		return nil
	}
	l := len(*flavors)
	for i := 0; i < l; i++ {
		// 添加套餐的id
		(*flavors)[i].SetmealID = dh.ID
	}
	// 插入套餐和菜品的关联表
	db.MealDishDao.InsertBatch(flavors)
	return nil
}
func PageQuerySetMeal(page *dto.SetmealPageQueryDTO) *common.PageResult {
	var pageResult = common.PageResult{}
	pageResult.Records, pageResult.Total = db.MealDao.PageQuery(page)
	return &pageResult
}
func DeleteBatchMeal(ids *[]int64) error {
	for i := 0; i < len(*ids); i++ {
		err := db.MealDao.GetByID((*ids)[i])
		//判断当前菜品是否能够删除---是否存在起售中的菜品？？
		if err.Status == status_c.ENABLE {
			//起售中的套餐不能删除
			hlog.Error(err)
			return nil
		}
	}
	for _, v := range *ids {
		db.MealDao.DeleteByID(v)
		db.MealDishDao.DeleteBySetmealId(v)
	}
	return nil
}
func GetByIdWithDishMeal(id int64) (*vo.SetmealVO, error) {
	v, err := db.MealDao.GtByIdWithDish(id)
	if err != nil {
		hlog.Infof(err.Error())
		return nil, err
	}
	return v, nil

}
func UpdateMeal(meal *model.Setmeal, dish *[]model.SetmealDish) {
	db.MealDao.Update(meal)
	db.MealDishDao.DeleteBySetmealId(meal.ID)
	id := meal.ID
	l := len(*dish)
	for i := 0; i < l; i++ {
		(*dish)[i].SetmealID = id
	}
	db.MealDishDao.InsertBatch(dish)
}
func StartOrStopMeal(status int32, id int64) error {
	//起售套餐时，判断套餐内是否有停售菜品，有停售菜品提示"套餐内包含未启售菜品，无法启售"
	if status == status_c.ENABLE {
		//select a.* from dish a left join setmeal_dish b on a.id = b.dish_id where b.setmeal_id = ?
		dishList := db.DisDao.GetBySetmealId(id)
		if dishList != nil && len(dishList) > 0 {
			for _, item := range dishList {
				if status_c.DISABLE == item.Status {
					return errors.New(message_c.SETMEAL_ENABLE_FAILED)
				}
			}
		}
	}
	setmeal := model.Setmeal{ID: id, Status: status}
	db.MealDao.Update(&setmeal)
	return nil
}
func ListSetmeal(meal *model.Setmeal) *[]model.Setmeal {
	return db.MealDao.List(meal)
}
