package db

import "reggie/internal/models/model"

type dishFI interface {
	InsertBatch(flavors *[]model.DishFlavor)
}
type dishFDao struct {
}

func (*dishFDao) InsertBatch(flavors *[]model.DishFlavor) {
	DBEngine.Select("*").Create(flavors)
}
