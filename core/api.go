package core

// APIInfo API 接口信息
type APIInfo struct {
	Method       string
	URI          string
	ResponseType string
}

// APIDomain 接口 domain
var APIDomain = "https://api.weixin.qq.com"

// WxAPI 接口列表
var WxAPI = map[string]map[string]APIInfo{
	"access_token": map[string]APIInfo{
		"get": APIInfo{
			Method:       "get",
			URI:          "/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET",
			ResponseType: "json",
		},
	},
	"oauth2": map[string]APIInfo{
		"access_token": APIInfo{
			Method:       "get",
			URI:          "/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code",
			ResponseType: "json",
		},
		"refresh_token": APIInfo{
			Method:       "get",
			URI:          "/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN",
			ResponseType: "json",
		},
		"auth": APIInfo{
			Method:       "get",
			URI:          "/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID",
			ResponseType: "json",
		},
		"userinfo": APIInfo{
			Method:       "get",
			URI:          "/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN",
			ResponseType: "json",
		},
	},
}
