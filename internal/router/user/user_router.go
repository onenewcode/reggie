package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
	"reggie/internal/models/common"
	"reggie/internal/models/dto"
	"reggie/internal/router/service"
)

func LoginUser(ctx context.Context, c *app.RequestContext) {
	var userLoginDto dto.UserLoginDTO
	c.Bind(&userLoginDto)
	hlog.Info("微信用户登录：{}", userLoginDto)
	us := service.WxLoginUser(&userLoginDto)
	c.JSON(http.StatusOK, common.Result{1, "", us})
}
