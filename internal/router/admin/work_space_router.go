package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
	"reggie/internal/models/common"
	"reggie/internal/router/service"
	"strconv"
	"time"
)

// @Summary 设置店铺的营业状态
func BusinessData(ctx context.Context, c *app.RequestContext) {
	now := time.Now()
	//获得当天的开始时间
	begin := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	//获得当天的结束时间
	end := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	c.JSON(http.StatusOK, common.Result{1, "", service.GetBusinessData(begin, end)})
}

// @Summary  查询订单管理数据
func OrderOverView(ctx context.Context, c *app.RequestContext) {

	c.JSON(http.StatusOK, common.Result{1, "", service.GetOrderOverView()})
}

// @Summary 设置店铺的营业状态
func DishOverView(ctx context.Context, c *app.RequestContext) {
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
func SetmealOverView(ctx context.Context, c *app.RequestContext) {
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
