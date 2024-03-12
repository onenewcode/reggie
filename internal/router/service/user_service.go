package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"reggie/internal/db"
	"reggie/internal/models/constant/message_c"
	"reggie/internal/models/dto"
	"reggie/internal/models/model"
	"reggie/pkg/wx"
)

func WxLoginUser(userLoginDTO *dto.UserLoginDTO) *model.User {
	op_id := wx.WxClient.GetOpenid(&userLoginDTO.Code)
	if op_id == nil {
		hlog.Error(message_c.LOGIN_FAILED)
	}
	var us *model.User
	us = db.UserDao.GetByOpenid(op_id)
	// 查询不到，就字段添加新用户
	if us == nil {
		us = db.UserDao.Insert(userLoginDTO.ToNewUser())
	}

	return us
}
