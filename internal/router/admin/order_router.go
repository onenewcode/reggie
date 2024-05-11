package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"reggie/internal/constant/message_c"
	"reggie/internal/dal/common"
	"reggie/internal/dal/dto"
	"reggie/internal/router/service"
)

// @Summary 订单搜索
func ConditionSearchOrder(ctx context.Context, c *app.RequestContext) {
	var oq dto.OrdersPageQueryDTO
	err := c.Bind(&oq)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Result{1, message_c.UNKNOWN_ERROR, nil})
	}
	r, err := service.ConditionSearchOrder(&oq)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Result{1, message_c.UNKNOWN_ERROR, nil})
	}
	c.JSON(http.StatusOK, r)
}
