package redis

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
	"reggie/internal/dal/vo"
	"strconv"
	"time"
)

const (
	dish_vo         = "dishVo:"
	windowSize      = 3 * time.Second // 滑动窗口大小
	maxRequests     = 10              // 最大请求数
	windowKeyPrefix = "sliding_window:"
)

type RedisClient interface {
	SetStatus(status *int)
	GetStatus() *int
	GetListDishVO(categoryId string) (*[]vo.DishVO, error)
	SetListDishVO(categoryId string, dvo *[]vo.DishVO) error
	ClearCacheDishByCategoryId(categoryId string)
	AllowRequest(clientId string) (bool, error)
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

// 使用滑动窗口限流
func (*redisClient) AllowRequest(user_id string) (bool, error) {
	ctx := context.Background()
	// 获取当前时间
	now := time.Now().UnixNano()
	before := now - windowSize.Nanoseconds()
	// 设置key
	key := windowKeyPrefix + user_id

	// 删除窗口之外的请求记录
	rc.ZRemRangeByScore(ctx, key, strconv.FormatInt(before, 10), strconv.FormatInt(now, 10))

	// 累加新请求
	pipe := rc.Pipeline() //批量发送命令提高吞吐量
	// 查询分数是否存在
	exists, _ := pipe.ZScore(ctx, key, strconv.FormatInt(now, 10)).Result()
	if exists == 0 {
		// 添加分数
		pipe.ZAdd(ctx, key, redis.Z{Score: float64(now), Member: now})
	} else {
		pipe.ZIncrBy(ctx, key, 1, strconv.FormatInt(now, 10))
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	// 检查请求是否超限
	count, err := rc.ZCard(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if count > int64(maxRequests) {
		return false, nil // 超限，拒绝请求
	}

	return true, nil // 未超限，允许请求
}
