package service

import (
	"reggie/internal/db"
	"reggie/internal/models/constant/status_c"
	"reggie/internal/models/vo"
	"time"
)

func GetBusinessData(begin time.Time, end time.Time) vo.BusinessDataVO {
	/**
	 * 营业额：当日已完成订单的总金额
	 * 有效订单：当日已完成订单的数量
	 * 订单完成率：有效订单数 / 总订单数
	 * 平均客单价：营业额 / 有效订单数
	 * 新增用户：当日新增用户的数量
	 */
	m := make(map[string]interface{})
	m["begin"] = begin
	m["end"] = end
	//查询总订单数
	totalOrderCount := db.OrderDao.CountByMap(m)
	m["status"] = status_c.COMPLETED
	//营业额
	turnover := db.OrderDao.SumByMap(m)
	//有效订单数
	validOrderCount := db.OrderDao.CountByMap(m)

	unitPrice, orderCompletionRate := 0.0, 0.0

	if totalOrderCount != 0 && validOrderCount != 0 {
		//订单完成率
		orderCompletionRate = float64(validOrderCount) / float64(totalOrderCount)
		//平均客单价
		unitPrice = turnover / float64(validOrderCount)
	}
	//新增用户数
	newUsers := db.UserDao.CountByMap(m)
	return vo.BusinessDataVO{
		Turnover:            turnover,
		ValidOrderCount:     validOrderCount,
		OrderCompletionRate: orderCompletionRate,
		UnitPrice:           unitPrice,
		NewUsers:            newUsers,
	}
}
func GetOrderOverView() vo.OrderOverViewVO {
	now := time.Now()
	begin := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	m := make(map[string]interface{})
	m["begin"] = begin
	m["status"] = status_c.TO_BE_CONFIRMED
	//待接单
	waitingOrders := db.OrderDao.CountByMap(m)
	//待派送
	m["status"] = status_c.CONFIRMED
	deliveredOrders := db.OrderDao.CountByMap(m)

	//已完成
	m["status"] = status_c.COMPLETED
	completedOrders := db.OrderDao.CountByMap(m)

	//已取消
	m["status"] = status_c.CANCELLED
	cancelledOrders := db.OrderDao.CountByMap(m)

	//全部订单
	m["status"] = nil
	allOrders := db.OrderDao.CountByMap(m)
	return vo.OrderOverViewVO{
		WaitingOrders:   waitingOrders,
		DeliveredOrders: deliveredOrders,
		CompletedOrders: completedOrders,
		CancelledOrders: cancelledOrders,
		AllOrders:       allOrders,
	}
}
func GetDishOverView() vo.DishOverViewVO {
	m := make(map[string]interface{})
	m["status"] = status_c.ENABLE
	sold := db.OrderDao.CountByMap(m)

	m["status"] = status_c.DISABLE
	discontinued := db.OrderDao.CountByMap(m)

	return vo.DishOverViewVO{
		Sold:         sold,
		Discontinued: discontinued,
	}
}
