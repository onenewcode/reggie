package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"reggie/internal/constant/message_c"
	"reggie/internal/dal/common"
	"reggie/internal/dal/dto"
	"reggie/internal/dal/model"
	"reggie/internal/middleware"
	"reggie/internal/router/service"
	"strconv"
	"time"
)

// over

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
		emp_id := middleware.GetJwtPayload(c)
		empL.CreateUser, empL.UpdateUser = emp_id, emp_id
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

// 禁用员工账号
// @Summary 禁用员工账号
// @Accept application/json
// @Produce application/json
// @router /admin/employee/status [post]
func StartOrStopEmp(ctx context.Context, c *app.RequestContext) {
	status, id := c.Param("status"), c.Query("id")
	log.Printf("启用禁用员工账号：{%s},{%s}", status, id)
	status_r, _ := strconv.ParseInt(status, 10, 32)
	id_r, _ := strconv.ParseInt(id, 10, 64)
	service.StartOrStopEmp(int32(status_r), id_r, middleware.GetJwtPayload(c))
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}

// 根据id查找雇员
// @Summary 根据id查找雇员
// @Accept application/json
// @Produce application/json
// @router /admin/employee/status [get]
func GetByIdEmp(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	log.Printf("查询员工账号：{%s}", id)
	id_r, _ := strconv.ParseInt(id, 10, 64)
	emp := service.GetByIdEmp(id_r)
	c.JSON(http.StatusOK, common.Result{1, "", emp})
}

// 更新雇员信息
// @Summary 根据id更新雇员信息
// @Accept application/json
// @Produce application/json
// @router /admin/employee [put]
func UpdateEmp(ctx context.Context, c *app.RequestContext) {
	var emp model.Employee
	c.BindAndValidate(&emp)
	emp.UpdateUser, emp.UpdateTime = middleware.GetJwtPayload(c), time.Now()
	log.Println("编辑员工信息：", emp)
	service.UpdateEmp(&emp)
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}
