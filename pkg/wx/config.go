package wx

const (
	wx_login = "https://api.weixin.qq.com/sns/jscode2session"
	app_id   = ""
	secret   = ""
)

var (
	WxClient wxInterface
)

func init() {
	WxClient = &wxClient{}
}
