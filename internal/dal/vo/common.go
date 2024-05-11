package vo

import (
	"reggie/internal/dal/model"
	"time"
)

type BusinessDataVO struct {
	Turnover float64 `json:"turnover,omitempty"` //营业额

	ValidOrderCount int64 `json:"valid_order_count,omitempty"` //有效订单数

	OrderCompletionRate float64 `json:"order_completion_rate,omitempty"` //订单完成率

	UnitPrice float64 `json:"unit_price,omitempty"` //平均客单价

	NewUsers int64 `json:"new_users,omitempty"` //新增用户数

}
type OrderOverViewVO struct {

	//待接单数量
	WaitingOrders int64 `json:"waiting_orders,omitempty"`

	//待派送数量
	DeliveredOrders int64 `json:"delivered_orders,omitempty"`

	//已完成数量
	CompletedOrders int64 `json:"completed_orders,omitempty"`

	//已取消数量
	CancelledOrders int64 `json:"cancelled_orders,omitempty"`

	//全部订单
	AllOrders int64 `json:"all_orders,omitempty"`
}

type EmployeeLoginVO struct {
	Id int64 `json:"id,omitempty"`

	UserName string `json:"user_name,omitempty"`

	Name string `json:"name,omitempty"`

	Token string `json:"token,omitempty"`
}
type DishOverViewVO struct {

	// 已启售数量
	Sold int64 `json:"sold,omitempty"`

	// 已停售数量
	Discontinued int64 `json:"discontinued,omitempty"`
}
type SetmealOverViewVO struct {

	// 已启售数量
	Sold int64 `json:"sold,omitempty"`

	// 已停售数量
	Discontinued int64 `json:"discontinued,omitempty"`
}

type SetmealVO struct {
	ID            int64               `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`  // 主键
	CategoryID    int64               `gorm:"column:category_id;not null;comment:菜品分类id" json:"category_id"` // 菜品分类id
	Name          string              `gorm:"column:name;not null;comment:套餐名称" json:"name"`                 // 套餐名称
	Price         float64             `gorm:"column:price;not null;comment:套餐价格" json:"price"`               // 套餐价格
	Status        int32               `gorm:"column:status;default:1;comment:售卖状态 0:停售 1:起售" json:"status"`  // 售卖状态 0:停售 1:起售
	Description   string              `gorm:"column:description;comment:描述信息" json:"description"`            // 描述信息
	Image         string              `gorm:"column:image;comment:图片" json:"image"`                          // 图片
	UpdateTime    time.Time           `gorm:"column:update_time;comment:更新时间" json:"update_time"`            // 更新时间
	SetmealDishes []model.SetmealDish `gorm:"foreignKey:setmeal_id"`
}
type UserLoginVO struct {
	Id     int64  `json:"id,omitempty"`
	Openid string `json:"openid,omitempty"`
	Token  string `json:"token,omitempty"`
}

func (ul *UserLoginVO) User2UserLoginVO(u *model.User) {
	ul.Id = (*u).ID
	ul.Openid = (*u).Openid
}
