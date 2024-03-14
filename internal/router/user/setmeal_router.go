package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"reggie/internal/models/common"
	"reggie/internal/models/constant/status_c"
	"reggie/internal/models/model"
	"reggie/internal/router/service"
	"strconv"
)

// 条件查询
func ListSetmeal(ctx context.Context, c *app.RequestContext) {
	ty_pe := c.Query("categoryId")
	log.Printf("按照类型查询菜品：{%s}", ty_pe)
	tp, _ := strconv.ParseInt(ty_pe, 10, 64)
	meal := model.Setmeal{
		CategoryID: tp,
		Status:     status_c.ENABLE,
	}
	new_meal := service.ListSetmeal(&meal)
	c.JSON(http.StatusOK, common.Result{1, "", new_meal})
}

// 根据套餐id查询包含的菜品列表
func DishListSetmeal(ctx context.Context, c *app.RequestContext) {
	id_s := c.Param("id")
	log.Printf("根据套餐id查询包含的菜品列表：{%s}", id_s)
	id, _ := strconv.ParseInt(id_s, 10, 64)
	meal, err := service.GetByIdWithDishMeal(id)
	if err != nil {
		c.JSON(http.StatusOK, common.Result{1, "", nil})
	}
	c.JSON(http.StatusOK, common.Result{1, "", meal})

}
