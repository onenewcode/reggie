package admin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"mime/multipart"
	"net/http"
	obs "reggie/internal/OBS"
	"reggie/internal/models/common"
	"reggie/internal/models/constant/message_c"
)

func getFile(from *multipart.Form) *multipart.FileHeader {
	fileH := from.File["file"][0]
	return fileH
}
func UploadImg(ctx context.Context, c *app.RequestContext) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK, common.Result{0, message_c.UPLOAD_FAILED, nil})
	}
	if str := obs.OBS.UploadImg(getFile(form)); str != nil {
		c.JSON(http.StatusOK, common.Result{1, "", str})
	}
}
