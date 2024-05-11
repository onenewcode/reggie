package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"mime/multipart"
	"net/http"
	"reggie/internal/constant/message_c"
	"reggie/internal/dal/common"
	"reggie/pkg/obs"
)

// over
func getFile(from *multipart.Form) *multipart.FileHeader {
	// 获取文件对应的文件头
	fileH := from.File["file"][0]
	return fileH
}
func UploadImg(ctx context.Context, c *app.RequestContext) {
	form, err := c.MultipartForm()
	hlog.Info("文件上传：{ %s}", form)
	if err != nil {
		c.JSON(http.StatusOK, common.Result{0, message_c.UPLOAD_FAILED, nil})
	}
	if str := obs.OBS.UploadImg(getFile(form)); str != nil {
		c.JSON(http.StatusOK, common.Result{1, "", str})
	}
}
