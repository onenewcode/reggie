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
	"reggie/pkg/redis"
	"strconv"
)

// 根据id查询菜品
func ListDish(ctx context.Context, c *app.RequestContext) {
	c_id := c.Query("categoryId")
	log.Printf("按照类型查询菜品：{%s}", c_id)
	id, _ := strconv.ParseInt(c_id, 10, 64)
	//查询redis中是否存在菜品数据
	list, err := redis.RC.GetListDishVO(c_id)
	// redis有数据直接返回
	if err == nil {
		c.JSON(http.StatusOK, common.Result{1, "", list})
	}
	// 无数据查询数据库
	d := model.Dish{
		CategoryID: id,
		Status:     status_c.ENABLE,
	}
	//如果不存在，查询数据库，将查询到的数据放入redis中
	list = service.ListWithFlavorDish(&d)
	//////////////////////////////////////////////////////////
	err = redis.RC.SetListDishVO(c_id, list)
	if err != nil {
		c.JSON(http.StatusOK, common.Result{0, "", list})
	}
	c.JSON(http.StatusOK, common.Result{1, "", list})

}
