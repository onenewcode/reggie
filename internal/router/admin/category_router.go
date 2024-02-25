package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"reggie/internal/middleware"
	"reggie/internal/models/common"
	"reggie/internal/models/model"
	"reggie/internal/router/service"
	"time"
)

// 新增菜品
// @Summary 新增菜品
// @Accept application/json
// @Produce application/json
// @router /admin/category [post]
func SaveCategory(ctx context.Context, c *app.RequestContext) {
	var category model.Category
	c.Bind(&category)
	// 赋予创建用户和更新用户的数据
	category.CreateUser, category.UpdateUser = middleware.GetJwtPayload(c), middleware.GetJwtPayload(c)
	// 赋予创建时间和更新时间数据
	category.CreateTime, category.UpdateTime = time.Now(), time.Now()
	log.Println("新增分类：", category)
	service.SaveCategory(&category)
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}
