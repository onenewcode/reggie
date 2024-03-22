package db

import (
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
	"time"
)

type orderI interface {
	Insert(order *model.Order) (*model.Order, error)
	PageQuery(page *dto.OrderPageQueryDTO) (*[]model.Order, error)
}
type orderDao struct {
}

func (*orderDao) Insert(order *model.Order) (*model.Order, error) {
	err := DBEngine.Create(order).Error
	if err != nil {
		return nil, err
	} else {
		return order, nil
	}
}
func (*orderDao) PageQuery(page *dto.OrderPageQueryDTO) (*[]model.Order, error) {
	var (
		order []model.Order
		//count int64
	)
	origin_sql := DBEngine
	// 判断是否含有name，有name不为nil，就进行模糊查询。
	if page.Number != "" {
		origin_sql = origin_sql.Where("number= ?", page.Number)
	}
	if page.Page != 0 {
		origin_sql = origin_sql.Where("phone  =?", page.Page)
	}
	if page.UserId != 0 {
		origin_sql = origin_sql.Where("userId=?", page.UserId)
	}
	if page.Status != 0 {
		origin_sql = origin_sql.Where("status =?", page.Status)
	}
	if page.BeginTime.Equal(time.Time{}) {
		origin_sql = origin_sql.Where("beginTime =?", page.BeginTime)
	}
	if page.EndTime.Equal(time.Time{}) {
		origin_sql = origin_sql.Where("endTime  =?", page.EndTime)
	}
	//origin_sql.Model(&model.Category{}).Count(&count)
	origin_sql.Find(&order)
	return &order, nil
}
