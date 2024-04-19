package db

import (
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
	"time"
)

type orderI interface {
	Insert(order *model.Order) (*model.Order, error)
	PageQuery(page *dto.OrdersPageQueryDTO) (*[]model.Order, int64, error)
	CountByMap(m map[string]interface{}) int64
	SumByMap(m map[string]interface{}) float64
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
func (*orderDao) PageQuery(page *dto.OrdersPageQueryDTO) (*[]model.Order, int64, error) {
	var (
		order []model.Order
		count int64
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
	origin_sql.Model(&model.Category{}).Count(&count)
	origin_sql.Find(&order)
	return &order, count, nil
}
func (*orderDao) CountByMap(m map[string]interface{}) int64 {
	var nums int64
	origin_sql := DBEngine
	// 判断是否含有name，有name不为nil，就进行模糊查询。
	if m["begin"] != nil {
		origin_sql = origin_sql.Where(" order_time >?", m["begin"])
	}
	if m["end"] != nil {
		origin_sql = origin_sql.Where(" order_time <?", m["end"])
	}
	if m["status"] != nil {
		origin_sql = origin_sql.Where("status=?", m["status"])
	}
	origin_sql.Table(model.TableNameOrder).Count(&nums)
	return nums
}
func (*orderDao) SumByMap(m map[string]interface{}) float64 {
	var nums float64
	origin_sql := DBEngine
	// 判断是否含有name，有name不为nil，就进行模糊查询。
	if m["begin"] != nil {
		origin_sql = origin_sql.Where(" order_time >?", m["begin"])
	}
	if m["end"] != nil {
		origin_sql = origin_sql.Where(" order_time <?", m["end"])
	}
	if m["status"] != nil {
		origin_sql = origin_sql.Where("status=?", m["status"])
	}
	origin_sql.Table(model.TableNameOrder).Select("SUM(price)").Scan(&nums)
	return nums
}
