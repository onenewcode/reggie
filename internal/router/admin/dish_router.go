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

//func SaveDish(ctx context.Context, c *app.RequestContext) {
//	var dish model.Dish
//	c.Bind(&dish)
//	// 赋予创建用户和更新用户的数据
//	dish.CreateUser, dish.UpdateUser = middleware.GetJwtPayload(c), middleware.GetJwtPayload(c)
//	// 赋予创建时间和更新时间数据
//	dish.CreateTime, dish.UpdateTime = time.Now(), time.Now()
//	log.Println("新增分类：", dish)
//	service.SaveDish(&dish)
//	c.JSON(http.StatusOK, common.Result{1, "", nil})
//}

func PageDish(ctx context.Context, c *app.RequestContext) {
	var dishPage dto.DishPageQueryDTO
	c.Bind(&dishPage)
	log.Println("菜品分类查询", dishPage)
	cat := service.PageQueryDish(&dishPage)
	c.JSON(http.StatusOK, common.Result{1, "", cat})
}

func DeleteDish(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	log.Printf("根据id删除菜品：{%s}", id)
	id_r, _ := strconv.ParseInt(id, 10, 64)
	if err := service.DeleteDish(&id_r); err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, common.Result{0, "", nil})
	} else {
		c.JSON(http.StatusOK, common.Result{1, "", nil})
	}
}

func UpdateDish(ctx context.Context, c *app.RequestContext) {
	var dish model.Dish
	c.Bind(&dish)
	// 赋予创建时间和更新时间数据
	dish.CreateTime, dish.UpdateTime = time.Now(), time.Now()
	log.Println("修改菜品分类信息：", dish)
	service.UpdateDish(&dish)
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}

func StartOrStopDish(ctx context.Context, c *app.RequestContext) {
	status, id := c.Param("status"), c.Query("id")
	log.Printf("启用禁用菜品分类：{%s},{%s}", status, id)
	status_r, _ := strconv.ParseInt(status, 10, 32)
	id_r, _ := strconv.ParseInt(id, 10, 64)
	service.StartOrStopDish(int32(status_r), id_r, middleware.GetJwtPayload(c))
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}

// 根据类型查询分类
func ListDish(ctx context.Context, c *app.RequestContext) {
	ty_pe := c.Query("type")
	log.Printf("按照类型查询菜品：{%s}", ty_pe)
	tp, _ := strconv.ParseInt(ty_pe, 10, 64)
	service.ListDish(&tp)
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}
func GetByIdDish(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	log.Printf("查询员工账号：{%s}", id)
	id_r, _ := strconv.ParseInt(id, 10, 64)
	emp := service.GetByIdEmp(id_r)
	c.JSON(http.StatusOK, common.Result{1, "", emp})
}
