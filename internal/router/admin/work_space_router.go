package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"reggie/internal/models/common"
	"reggie/internal/router/service"
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

// @Summary 查询菜品总览
func DishOverView(ctx context.Context, c *app.RequestContext) {

	c.JSON(http.StatusOK, common.Result{1, "", service.GetDishOverView()})
}

// @Summary 查询套餐总览
func SetmealOverView(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, common.Result{1, "", service.GetSetmealOverView()})
}
