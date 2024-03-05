package redis

import "context"

type RedisClient interface {
	SetStatus(status *int)
	GetStatus() *int
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
