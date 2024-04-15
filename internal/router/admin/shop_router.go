package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
	"reggie/internal/models/common"
	"reggie/internal/router/service"
	"strconv"
)

// @Summary 设置店铺的营业状态
func SetStatusShop(ctx context.Context, c *app.RequestContext) {
	s := c.Param("status")
	status, _ := strconv.Atoi(s)
	var statusString string
	if status == 1 {
		statusString = "营业中"
	} else {
		statusString = "打烊中"
	}
	hlog.Infof("设置店铺的营业状态为：", statusString)
	service.SetStatusShop(&status)
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}

// @Summary 获取店铺的营业状态
func GetStatusShop(ctx context.Context, c *app.RequestContext) {
	status := *service.GetStatusShop()
	var statusString string
	if status == 1 {
		statusString = "营业中"
	} else {
		statusString = "打烊中"
	}
	hlog.Infof("获取到店铺的营业状态为：{}", statusString)
	c.JSON(http.StatusOK, common.Result{1, "", status})
}
