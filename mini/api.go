package mini

import (
	"net/http"

	"github.com/zjxpcyc/wechat/core"
)

// API 接口列表
var API = map[string]map[string]core.API{
	"oauth2": map[string]core.API{
		"session": core.API{
			Method:       http.MethodGet,
			URI:          "https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code",
			ResponseType: "json",
		},
	},
}
