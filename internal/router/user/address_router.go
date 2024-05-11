package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
	"reggie/internal/constant/message_c"
	"reggie/internal/dal/common"
	"reggie/internal/dal/model"
	"reggie/internal/middleware"
	"reggie/internal/router/service"
	"strconv"
)

//	查询当前登录用户的所有地址信息
//
// list
func ListAddress(ctx context.Context, c *app.RequestContext) {
	address := model.AddressBook{}
	address.UserID = middleware.GetJwtPayload(c)
	list := service.ListAddress(&address)
	c.JSON(http.StatusOK, common.Result{1, "", list})
}

// 根据id查询地址
func SaveAddress(ctx context.Context, c *app.RequestContext) {
	address := model.AddressBook{}
	c.Bind(&address)
	address.UserID = middleware.GetJwtPayload(c)
	service.SaveAddress(&address)
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}

// @GetMapping("/{id}")
func GetByIdAddress(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	hlog.Infof("根据id查询地址：", id)
	id_r, _ := strconv.ParseInt(id, 10, 64)
	v, err := service.GetByIdAddress(id_r)
	if err != nil {
		c.JSON(http.StatusNotFound, common.Result{0, message_c.UNKNOWN_ERROR, nil})
	} else {
		c.JSON(http.StatusOK, common.Result{1, "", v})
	}
}

// 根据id修改地址
func UpdateAddress(ctx context.Context, c *app.RequestContext) {
	address := model.AddressBook{}
	c.Bind(&address)
	address.UserID = middleware.GetJwtPayload(c)
	service.UpdateAddress(&address)
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}

// 设置默认地址
func SetDefaultAddress(ctx context.Context, c *app.RequestContext) {
	address := model.AddressBook{}
	c.Bind(&address)
	address.UserID = middleware.GetJwtPayload(c)
	service.SetDefaultAddress(&address)
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}

// 根据id删除地址
func DeleteByIdAddress(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	hlog.Infof("根据id查询地址：", id)
	id_r, _ := strconv.ParseInt(id, 10, 64)
	service.DeleteByIdAddress(id_r)
	c.JSON(http.StatusOK, common.Result{1, "", nil})
}

//  查询默认地址

// @GetMapping("default")
// @ApiOperation("查询默认地址")
func GetDefaultAddress(ctx context.Context, c *app.RequestContext) {
	address := model.AddressBook{}
	address.UserID = middleware.GetJwtPayload(c)
	address.IsDefault = true
	list := service.ListAddress(&address)
	if len(*list) == 1 {
		c.JSON(http.StatusOK, common.Result{1, "", (*list)[0]})
	} else {
		c.JSON(http.StatusNotFound, common.Result{1, "没有查询到默认地址", nil})
	}

}
