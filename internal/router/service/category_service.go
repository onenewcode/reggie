package service

import (
	"reggie/internal/dal/common"
	"reggie/internal/dal/dto"
	"reggie/internal/dal/model"
	"reggie/internal/db"
	"time"
)

func SaveCategory(category *model.Category) {
	db.CatDao.Save(category)
}
func PageQueryDat(categoryPage *dto.CategoryPageQueryDTO) *common.PageResult {
	var pageResult = common.PageResult{}
	pageResult.Records, pageResult.Total = db.CatDao.PageQuery(categoryPage)
	return &pageResult
}
func DeleteCat(id *int64) *error {
	err := db.CatDao.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
func UpdateCategory(cat *model.Category) {
	db.CatDao.Update(cat)
}
func StartOrStopCat(status int32, id int64, update_user int64) {
	cat := model.Category{
		ID:         id,
		Status:     status,
		UpdateUser: update_user,
		UpdateTime: time.Now(),
	}
	db.CatDao.UpdateStatus(&cat)
}
func ListCat(tp *int64) *[]model.Category {
	return db.CatDao.List(tp)
}
