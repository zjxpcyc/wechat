package mini

import (
	"github.com/zjxpcyc/wechat/core"
)

// WxAPI 接口列表
var WxAPI = map[string]map[string]core.APIInfo{
	"oauth2": map[string]core.APIInfo{
		"session": core.APIInfo{
			Method:       "get",
			URI:          "https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code",
			ResponseType: "json",
		},
	},
}
