package service

import (
	"reggie/internal/db"
	"reggie/internal/models/common"
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
)

func SaveCategory(category *model.Category) {
	db.CatDao.Save(category)
}
func PageQueryDat(categoryPage *dto.CategoryPageQueryDTO) *common.PageResult {
	var pageResult = common.PageResult{}
	pageResult.Records, pageResult.Total = db.CatDao.PageQuery(categoryPage)
	return &pageResult
}
