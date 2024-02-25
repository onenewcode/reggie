package service

import (
	"reggie/internal/db"
	"reggie/internal/models/model"
)

func SaveCategory(category *model.Category) {
	db.CatDao.Save(category)
}
