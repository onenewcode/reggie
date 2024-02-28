package service

import (
	"reggie/internal/db"
	"reggie/internal/models/common"
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
	"time"
)

func SaveDish(category *model.Category) {
	db.CatDao.Save(category)
}
func PageQueryDish(categoryPage *dto.DishPageQueryDTO) *common.PageResult {
	var pageResult = common.PageResult{}
	pageResult.Records, pageResult.Total = db.DisDao.PageQuery(categoryPage)
	return &pageResult
}
func DeleteDish(id *int64) *error {
	err := db.DisDao.Delete(id)
	if err != nil {
		return err
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
