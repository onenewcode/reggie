package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"log"
	"net/http"
	"reggie/internal/middleware"
	"reggie/internal/models/common"
	"reggie/internal/models/constant/message_c"
	"reggie/internal/models/dto"
	"reggie/internal/router/service"
	"strconv"
	"strings"
)

// @Summary 套餐管理接机口
func SaveSetMealWithDish(ctx context.Context, c *app.RequestContext) {
	var setD dto.SetmealDTO
	// 参数绑定转化为结构体
	err := c.Bind(&setD)
	if err != nil {
		hlog.Error("新增套餐错误")
	}
	log.Printf("新增套餐:{%s}", setD)
	meal := setD.ToNewSetmeal()
	meal.CreateUser, meal.UpdateUser = middleware.GetJwtPayload(c), middleware.GetJwtPayload(c)
	err = service.SaveSetMealWithDish(meal, setD.SetmealDishes)
	if err == nil {
		c.JSON(http.StatusOK, common.Result{0, message_c.ALREADY_EXISTS, nil})
	}
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}

// @Summary 分页查询
func PageSetMeal(ctx context.Context, c *app.RequestContext) {
	var page dto.SetmealPageQueryDTO
	// 参数绑定转化为结构体
	c.Bind(&page)

	log.Println("套餐分页查询，参数为：", page.Name)

	c.JSON(http.StatusOK, common.Result{1, "", service.PageQuerySetMeal(&page)})
}

// @Summary 批量删除套餐批量删除套餐
func DeleteBatchMeal(ctx context.Context, c *app.RequestContext) {
	id := c.Query("ids")
	nums := make([]int64, 0, 5)
	log.Printf("根据id删除菜品：{%s}", id)
	ids := strings.Split(id, ",")
	// 转换成数字数组
	for _, v := range ids {
		id_r, _ := strconv.ParseInt(v, 10, 64)
		nums = append(nums, id_r)
	}

	if err := service.DeleteBatchMeal(&nums); err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, common.Result{0, "", nil})
	} else {
		c.JSON(http.StatusOK, common.Result{1, "", nil})
	}
}

// @Summary 根据id查询套餐
func GetByIDDishMeal(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	log.Printf("查询菜单：{%s}", id)
	id_r, _ := strconv.ParseInt(id, 10, 64)
	v, err := service.GetByIdWithDishMeal(id_r)
	if err != nil {
		c.JSON(http.StatusNotFound, common.Result{0, message_c.UNKNOWN_ERROR, nil})
	} else {
		c.JSON(http.StatusOK, common.Result{1, "", v})
	}

}

// @Summary 修改套餐
func UpdateMeal(ctx context.Context, c *app.RequestContext) {
	var meal_t dto.SetmealDTO
	c.Bind(&meal_t)
	// 赋予创建时间和更新时间数据
	meal := meal_t.ToNewSetmeal()
	meal.CreateUser, meal.CreateUser = middleware.GetJwtPayload(c), middleware.GetJwtPayload(c)
	log.Println("修改菜品分类信息：", meal)
	service.UpdateMeal(meal, meal_t.SetmealDishes)
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}

// 待完善
func StartOrStopMeal(ctx context.Context, c *app.RequestContext) {
	status, id := c.Param("status"), c.Query("id")
	log.Printf("启用禁用套餐分类：{%s},{%s}", status, id)
	status_r, _ := strconv.ParseInt(status, 10, 32)
	id_r, _ := strconv.ParseInt(id, 10, 64)
	service.StartOrStopMeal(int32(status_r), id_r, middleware.GetJwtPayload(c))
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}
