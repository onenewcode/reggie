package wx

const (
	wx_login = "https://api.weixin.qq.com/sns/jscode2session?"
	app_id   = "wx7c324ecaed51956d"
	secret   = "75654f62622687f1ef7a484a1d96d614"
)

var (
	WxClient wxInterface
)

func init() {
	WxClient = &wxClient{}
}
