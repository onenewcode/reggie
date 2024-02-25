package db

import (
	"reggie/internal/models/model"
)

type CategoryDao struct {
}

func (*CategoryDao) Save(category *model.Category) {
	DBEngine.Create(category)
}
