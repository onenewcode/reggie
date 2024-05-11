package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"reggie/internal/dal/common"
	"reggie/internal/router/service"
	"strconv"
)

// 分类查询
func ListCategory(ctx context.Context, c *app.RequestContext) {
	ty_pe := c.Query("type")
	log.Printf("按照类型查询菜品：{%s}", ty_pe)
	tp, _ := strconv.ParseInt(ty_pe, 10, 64)
	dish := service.ListCat(&tp)
	c.JSON(http.StatusOK, common.Result{1, "", dish})
}
