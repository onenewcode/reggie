package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"reggie/internal/db"
	"reggie/internal/models/common"
	"reggie/internal/models/constant/message_c"
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
	"reggie/internal/models/vo"
	"strconv"
	"time"
)

func SubmitOrder(order *model.Order) (*vo.OrderSubmitVO, error) {
	//1. 处理各种业务异常（地址簿为空、购物车数据为空）
	addressBook := db.AddressDA0.GetById(order.AddressBookID)
	if addressBook == nil {
		return nil, errors.New(message_c.ADDRESS_BOOK_IS_NULL)
	}
	shoppingCart := model.ShoppingCart{
		UserID: order.UserID,
	}
	shoppingCartList := db.ShopCartDao.List(&shoppingCart)
	list_len := len(*shoppingCartList)
	if list_len == 0 {
		return nil, errors.New(message_c.SHOPPING_CART_IS_NULL)
	}

	//2. 向订单表插入1条数据
	order.OrderTime = time.Now()
	order.PayStatus = model.UN_PAID
	order.Status = model.PENDING_PAYMENT
	order.Number = strconv.FormatInt(time.Now().UnixNano(), 10)
	order.Address = addressBook.Detail
	order.Phone = addressBook.Consignee
	//orders.setUserId(userId);

	db.OrderDao.Insert(order)
	orderDetailList := make([]model.OrderDetail, 0, list_len)
	//3. 向订单明细表插入n条数据
	for i := 0; i < list_len; i++ {
		//订单明细
		orderDetail := model.OrderDetail{}
		err := copier.Copy(&orderDetail, &(*shoppingCartList)[i])
		if err != nil {
			return nil, errors.New("类型赋值错误")
		}
		orderDetail.OrderID = order.ID
		orderDetailList[i] = orderDetail
	}
	db.OrderDetailDao.InsertBatch(&orderDetailList)
	//4. 清空当前用户的购物车数据
	db.ShopCartDao.DeleteByUserId(order.UserID)
	//5. 封装VO返回结果
	orderSubmitVO := vo.OrderSubmitVO{}
	err := copier.Copy(&orderSubmitVO, &order)
	if err != nil {
		return nil, errors.New("类型赋值错误")
	}
	return &orderSubmitVO, nil
}

func PageQuery4UserOrder(page *dto.OrdersPageQueryDTO) (common.PageResult, error) {
	// 分页条件查询
	query, err := db.OrderDao.PageQuery(page)
	if err != nil {
		return common.PageResult{}, err
	}
	list := make([]vo.OrderVO, 0, 10)
	// 查询出订单明细，并封装入OrderVO进行响应
	l := len(*query)
	if l > 0 {
		for i := 0; i < l; i++ {
			orderDetails := db.OrderDetailDao.GetByOrderId((*query)[i].UserID)
			orderVO := vo.OrderVO{}
			copier.Copy(&orderVO, orderDetails)
			orderVO.OrderDetailList = orderDetails
			list[i] = orderVO
		}
	}

	return common.PageResult{int64(l), list}, nil
}
func ConditionSearchOrder() {

}
