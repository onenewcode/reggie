package db

import "reggie/internal/dal/model"

type order_detialI interface {
	InsertBatch(list *[]model.OrderDetail) error
	GetByOrderId(id int64) *[]model.OrderDetail
}
type order_detialDao struct {
}

func (*order_detialDao) InsertBatch(list *[]model.OrderDetail) error {
	err := DBEngine.Create(list).Error
	if err != nil {
		return err
	}
	return nil
}
func (*order_detialDao) GetByOrderId(id int64) *[]model.OrderDetail {
	var list []model.OrderDetail
	DBEngine.Where("order_id = ?", id).Find(&list)
	return &list
}
