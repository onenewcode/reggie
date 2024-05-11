package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
	"reggie/internal/dal/common"
	"reggie/internal/dal/dto"
	"reggie/internal/middleware"
	"reggie/internal/router/service"
)

// 添加购物车
// /user/shoppingCart/add
func AddShoppingCart(ctx context.Context, c *app.RequestContext) {
	var shop_dto dto.ShoppingCartDTO
	c.Bind(&shop_dto)
	shop_cart := shop_dto.ToShoppingCart()
	shop_cart.UserID = middleware.GetJwtPayload(c)
	hlog.Infof("添加购物车：", shop_cart)
	service.AddShoppingCart(shop_cart)
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}
func ListShoppingCart(ctx context.Context, c *app.RequestContext) {
	// user
	user_id := middleware.GetJwtPayload(c)
	c.JSON(http.StatusOK, common.Result{1, "", service.ShowShoppingCart(user_id)})
}

// 清空购物车
func CleanShoppingCart(ctx context.Context, c *app.RequestContext) {
	user_id := middleware.GetJwtPayload(c)
	// 获取user_id
	service.CleanShoppingCart(user_id)
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}
