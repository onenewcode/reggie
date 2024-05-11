package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"reggie/internal/dal/common"
	"reggie/internal/dal/model"
	"reggie/internal/middleware"
	"reggie/internal/models/dto"
	"reggie/internal/router/service"
)

func SubmitOrders(ctx context.Context, c *app.RequestContext) {
	var order model.Order
	c.Bind(&order)
	order.UserID = middleware.GetJwtPayload(c)
	service.SubmitOrder(&order)
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}
func PageOrders(ctx context.Context, c *app.RequestContext) {
	var page dto.OrderPageQueryDTO
	c.Bind(&page)
	page.UserId = middleware.GetJwtPayload(c)
	service.PageQuery4UserOrder(&page)
}
func DetailsOrders(ctx context.Context, c *app.RequestContext) {

}
func CancelOrders(ctx context.Context, c *app.RequestContext) {

}
