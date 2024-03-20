package redis

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"reggie/internal/models/vo"
)

const (
	dish_vo = "dishVo:"
)

type RedisClient interface {
	SetStatus(status *int)
	GetStatus() *int
	GetListDishVO(categoryId string) (*[]vo.DishVO, error)
	SetListDishVO(categoryId string, dvo *[]vo.DishVO) error
	ClearCacheDishByCategoryId(categoryId string)
}
type redisClient struct {
}

func (*redisClient) SetStatus(status *int) {
	rc.Set(
		context.Background(),
		shop_key,
		status,
		0,
	)
}
func (*redisClient) GetStatus() *int {
	val, _ := rc.Get(context.Background(), shop_key).Int()
	return &val
}

// 获取失败直接返回nil
func (*redisClient) GetListDishVO(categoryId string) (*[]vo.DishVO, error) {
	var dvo []vo.DishVO
	b, err := rc.Get(context.Background(), dish_vo+categoryId).Bytes()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &dvo); err != nil {
		hlog.Error("redis 解析ListDishVO失败")
		return nil, err
	}
	return &dvo, nil
}
func (*redisClient) SetListDishVO(categoryId string, dvo *[]vo.DishVO) error {
	b, err := json.Marshal(dvo)
	if err != nil {
		hlog.Error("redis 编码ListDishVO失败")
		return err
	}
	err = rc.Set(context.Background(), dish_vo+categoryId, b, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
func (*redisClient) ClearCacheDishByCategoryId(categoryId string) {
	rc.Del(context.Background(), dish_vo+categoryId)
}
