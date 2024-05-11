package status_c

/**
 * 状态常量，启用或者禁用
 */
const (
	//启用
	ENABLE int32 = 1

	//禁用
	DISABLE int32 = 0
	ALL     int32 = 3
	//设置新用户的默认密码
	DEFAULT_PASSWORD = "123456"
	//启用
)
const (
	/**
	 * 订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消
	 */
	PENDING_PAYMENT = iota + 1
	TO_BE_CONFIRMED
	CONFIRMED
	DELIVERY_IN_PROGRESS
	COMPLETED
	CANCELLED
)
