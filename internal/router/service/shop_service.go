package service

import "reggie/pkg/redis"

func SetStatusShop(status *int) {
	redis.RC.SetStatus(status)
}
func GetStatusShop() *int {
	return redis.RC.GetStatus()
}
