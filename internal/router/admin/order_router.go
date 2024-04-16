package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"reggie/internal/models/dto"
)

// @Summary 订单搜索
func ConditionSearchOrder(ctx context.Context, c *app.RequestContext) {
	var oq dto.OrdersPageQueryDTO
	c.Bind(&oq)

}
