package wx

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reggie/internal/models/dto"
)

type wxInterface interface {
	GetOpenid(code *string) *string
}
type wxClient struct {
}

// 获取opendid，
func (*wxClient) GetOpenid(code *string) *string {
	params := url.Values{}
	params.Add("appid", app_id)
	params.Add("secret", secret)
	params.Add("js_code", *code)
	params.Add("grant_type", "authorization_code")

	// 构建完整的URL
	fullURL := wx_login + params.Encode()

	resp, err := http.Get(fullURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var wxl dto.WXLoginDto
	json.Unmarshal(body, &wxl)

	return &wxl.OpenID
}
