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

var (
	empDao = db.EmployeeDao{}
)

func Login(ctx context.Context, c *app.RequestContext) {
	var empL model.Employee
	// 参数绑定转化为结构体
	err := c.Bind(&empL)
	if err != nil {
		log.Println("Employee 参数绑定失败")
	}
	//password := c.Query("password")
	emp := empDao.GetByUserName(empL.Username)
	c.JSON(http.StatusOK, common.Result{1, "2", emp})

}
