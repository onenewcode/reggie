package vo

import (
	"encoding/json"
	"reggie/internal/dal/model"
	"time"
)

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

func (dv *DishVO) ForDishAndFlavor(d *model.Dish, f *[]model.DishFlavor) {
	dv.Id, dv.Name, dv.CategoryId, dv.Price, dv.Image, dv.Description, dv.Status, dv.UpdateTime, dv.Flavors =
		d.ID, d.Name, d.CategoryID, d.Price, d.Image, d.Description, d.Status, d.UpdateTime, f
}
func Dish2DishVO(d *model.Dish) *DishVO {
	var dv DishVO
	dv.Id, dv.Name, dv.CategoryId, dv.Price, dv.Image, dv.Description, dv.Status, dv.UpdateTime = d.ID, d.Name, d.CategoryID, d.Price, d.Image, d.Description, d.Status, d.UpdateTime
	return &dv
}

// 添加josn序列化，方便redis存储
func (dv *DishVO) MarshalBinary() ([]byte, error) {
	return json.Marshal(dv)
}
func (dv *DishVO) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, dv)
}
