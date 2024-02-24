package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"log"
	"net/http"
	"reggie/internal/middleware"
	"reggie/internal/models/common"
	"reggie/internal/models/constant/message_c"
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
	"reggie/internal/router/service"
)

// 存储用户
// @Summary 存储用户
// @Accept application/json
// @Produce application/json
// @router /admin/employee [post]
func SaveEmp(ctx context.Context, c *app.RequestContext) {
	var empL model.Employee
	// 参数绑定转化为结构体
	err := c.Bind(&empL)
	if err != nil {
		log.Println("Employee 参数绑定失败")
		c.JSON(http.StatusBadRequest, common.Result{1, message_c.UNKNOWN_ERROR, nil})
	} else {
		// 获取jwt_payload的信息,并把信息赋予empL
		{
			jwt_payload, _ := c.Get("JWT_PAYLOAD")
			// 类型转换,我们的数据在claims中是以map[string]interface{}嵌套结构组成的。
			claims := jwt_payload.(jwt.MapClaims)
			origin_emp := claims[middleware.IdentityKey].(map[string]interface{})
			emp_id := origin_emp["id"].(float64)
			empL.CreateUser, empL.UpdateUser = int64(emp_id), int64(emp_id)
		}
		log.Printf("新增用户:{%s}", empL.Username)
		flag := service.SaveEmp(&empL)
		if flag == true {
			c.JSON(http.StatusOK, common.Result{1, "", nil})
		}
		c.JSON(http.StatusBadRequest, common.Result{1, message_c.ALREADY_EXISTS, nil})
	}
}

// 分页查询
// @Summary 分页查询
// @Accept application/json
// @Produce application/json
// @router /admin/employee/page [get]
func PageEmp(ctx context.Context, c *app.RequestContext) {
	var page dto.EmployeePageQueryDTO
	// 参数绑定转化为结构体
	err := c.Bind(&page)
	if err != nil {
		log.Println("Employee 参数绑定失败")
		c.JSON(http.StatusBadRequest, common.Result{1, message_c.UNKNOWN_ERROR, nil})
	} else {
		log.Println("员工分页查询，参数为：", page.Name)

		c.JSON(http.StatusOK, common.Result{1, "", service.PageQueryEmp(&page)})
	}
}
