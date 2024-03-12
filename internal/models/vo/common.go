package vo

import (
	"reggie/internal/models/model"
	"time"
)

type EmployeeLoginVO struct {
	Id int64 `json:"id,omitempty"`

	UserName string `json:"user_name,omitempty"`

	Name string `json:"name,omitempty"`

	Token string `json:"token,omitempty"`
}
type DishVO struct {
	Id int64 `json:"id,omitempty"`
	//菜品名称
	Name string `json:"name,omitempty"`
	//菜品分类id
	CategoryId int64 `json:"category_id,omitempty"`
	//菜品价格
	Price float64 `json:"price,omitempty"`
	//图片
	Image string `json:"image,omitempty"`
	//描述信息
	Description string `json:"description,omitempty"`
	//0 停售 1 起售
	Status int32 `json:"status,omitempty"`
	//更新时间
	UpdateTime time.Time `json:"update_time"`
	//分类名称
	CategoryName string `json:"category_name,omitempty"`
	//菜品关联的口味
	Flavors *[]model.DishFlavor `json:"flavors,omitempty"`
}

func (dt *DishVO) ForDishAndFlavor(d *model.Dish, f *[]model.DishFlavor) {
	dt.Id, dt.Name, dt.CategoryId, dt.Price, dt.Image, dt.Description, dt.Status, dt.UpdateTime, dt.Flavors =
		d.ID, d.Name, d.CategoryID, d.Price, d.Image, d.Description, d.Status, d.UpdateTime, f
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
