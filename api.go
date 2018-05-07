package wechat

import (
	"github.com/zjxpcyc/wechat/core"
)

// WxAPI 接口列表
var WxAPI = map[string]map[string]core.APIInfo{
	"access_token": map[string]core.APIInfo{
		"get": core.APIInfo{
			Method:       "get",
			URI:          "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET",
			ResponseType: "json",
		},
	},
	"oauth2": map[string]core.APIInfo{
		"access_token": core.APIInfo{
			Method:       "get",
			URI:          "https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code",
			ResponseType: "json",
		},
		"refresh_token": core.APIInfo{
			Method:       "get",
			URI:          "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN",
			ResponseType: "json",
		},
		"auth": core.APIInfo{
			Method:       "get",
			URI:          "https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID",
			ResponseType: "json",
		},
		"userinfo": core.APIInfo{
			Method:       "get",
			URI:          "https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN",
			ResponseType: "json",
		},
	},
}
