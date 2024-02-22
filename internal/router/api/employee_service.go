package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"reggie/internal/db"
	"reggie/internal/models/common"
	"reggie/internal/models/model"
)

// PingHandler 测试 handler
// @Summary 测试 Summary
// @Description 测试 Description
// @Accept application/json
// @Produce application/json
// @Router /ping [get]
func Login(ctx context.Context, c *app.RequestContext) {
	var empL model.Employee
	// 参数绑定转化为结构体
	err := c.Bind(&empL)
	if err != nil {
		log.Println("Employee 参数绑定失败")
	}

	//password := c.Query("password")
	emp := db.EmpDao.GetByUserName(empL.Username)
	if emp == nil {
		c.JSON(http.StatusNotFound, common.Result{0, "未知用户", nil})
	}
	c.JSON(http.StatusOK, common.Result{1, "", emp})

}
