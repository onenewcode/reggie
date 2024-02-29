package service

import (
	"reggie/internal/db"
	"reggie/internal/models/common"
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
	"time"
)

func SaveDish(dish *model.Dish) {
	db.DisDao.Save(dish)
}
func PageQueryDish(categoryPage *dto.DishPageQueryDTO) *common.PageResult {
	var pageResult = common.PageResult{}
	pageResult.Records, pageResult.Total = db.DisDao.PageQuery(categoryPage)
	return &pageResult
}
func DeleteDish(ids *[]int64) *error {
	for i := 0; i < len(*ids); i++ {

		err := db.DisDao.Delete((*ids)[i])
		if err != nil {
			return err
		}
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
func GetByIdDish(id int64) *model.Dish {
	return db.DisDao.GetById(id)
}
