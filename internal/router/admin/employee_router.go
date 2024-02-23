package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"reggie/internal/models/common"
	"reggie/internal/models/constant/message_c"
	"reggie/internal/models/model"
	"reggie/internal/router/service"
)

// 存储用户
// @Summary 存储用户
// @Accept application/json
// @Produce application/json
// @router /admin/employee [post]
func Save(ctx context.Context, c *app.RequestContext) {
	var empL model.Employee
	// 参数绑定转化为结构体
	err := c.Bind(&empL)
	if err != nil {
		log.Println("Employee 参数绑定失败")
	}
	log.Printf("新增用户:{%s}", empL.Username)
	flag := service.SavEmp(&empL)
	if flag == true {
		c.JSON(http.StatusOK, common.Result{1, "", nil})
	}
	c.JSON(http.StatusBadRequest, common.Result{1, message_c.ALREADY_EXISTS, nil})
}
