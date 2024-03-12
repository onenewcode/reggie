package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"reggie/internal/db"
	"reggie/internal/models/common"
	"reggie/internal/models/constant/message_c"
	"reggie/internal/models/constant/status_c"
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
	"reggie/internal/models/vo"
	"time"
)

// 第一个参数菜品，第二个参数菜品口味数组
func SaveWithFlavorDish(dish *model.Dish, flavors *[]model.DishFlavor) *model.Dish {
	db.DisDao.Save(dish)
	for _, v := range *flavors {
		v.DishID = dish.ID
	}
	db.DishFDao.InsertBatch(flavors)
	return dish
}
func PageQueryDish(categoryPage *dto.DishPageQueryDTO) *common.PageResult {
	var pageResult = common.PageResult{}
	pageResult.Records, pageResult.Total = db.DisDao.PageQuery(categoryPage)
	return &pageResult
}
func DeleteDish(ids *[]int64) *error {
	for i := 0; i < len(*ids); i++ {
		err := db.DisDao.GetById((*ids)[i])
		//判断当前菜品是否能够删除---是否存在起售中的菜品？？
		if err.Status == status_c.ENABLE {
			//当前菜品处于起售中，不能删除
			hlog.Error(err)
			return nil
		}
	}
	//判断当前菜品是否能够删除---是否被套餐关联了？？
	nums := db.MealDishDao.GetSetmealIdsByDishIds(ids)
	if len(*nums) != 0 {
		//当前菜品被套餐关联了，不能删除
		hlog.Error(message_c.DISH_BE_RELATED_BY_SETMEAL)
		return nil
	}
	//删除菜品表中的菜品数据
	for i := 0; i < len(*ids); i++ {
		db.DisDao.Delete((*ids)[i])
		//删除菜品关联的口味数据
		db.DishFDao.DeleteByDishId((*ids)[i])
	}
	return nil
}
func UpdateDish(dish *model.Dish) {
	db.DisDao.Update(dish)
}
func StartOrStopDish(status int32, id int64, update_user int64) {
	cat := model.Dish{
		ID:         id,
		Status:     status,
		UpdateUser: update_user,
		UpdateTime: time.Now(),
	}
	db.DisDao.UpdateStatus(&cat)
}
func ListDish(tp *int64) *[]model.Dish {
	return db.DisDao.List(tp)
}
func GetByIdWithFlavor(id int64) *vo.DishVO {
	dvo := &vo.DishVO{}
	//根据id查询菜品数据
	dish := db.DisDao.GetById(id)

	//根据菜品id查询口味数据

	dishFlavors := db.DishFDao.GetByDishId(id)
	dvo.ForDishAndFlavor(dish, dishFlavors)
	return dvo
}
