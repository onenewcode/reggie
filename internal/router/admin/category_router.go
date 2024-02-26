package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"reggie/internal/middleware"
	"reggie/internal/models/common"
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
	"reggie/internal/router/service"
	"strconv"
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

// 菜品分类分页
// @Summary 新增菜品
// @Accept application/json
// @Produce application/json
// @router /admin/category/page [get]
func PageCat(ctx context.Context, c *app.RequestContext) {
	var categoryPage dto.CategoryPageQueryDTO
	c.Bind(&categoryPage)
	log.Println("菜品分类查询", categoryPage)
	cat := service.PageQueryDat(&categoryPage)
	c.JSON(http.StatusOK, common.Result{1, "", cat})
}

// 按照id删除菜品分类
// @Summary 新增菜品
// @Accept application/json
// @Produce application/json
// @router /admin/category [delete]
func DeleteCat(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	log.Printf("查询员工账号：{%s}", id)
	id_r, _ := strconv.ParseInt(id, 10, 64)
	if err := service.DeleteCat(&id_r); err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, common.Result{0, "", nil})
	} else {
		c.JSON(http.StatusOK, common.Result{1, "", nil})
	}

}
