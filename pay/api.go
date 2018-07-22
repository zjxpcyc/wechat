package pay

import (
	"net/http"

	"github.com/zjxpcyc/wechat/core"
)

// API 接口列表
var API = map[string]map[string]core.API{
	"order": map[string]core.API{
		"post": core.API{
			Method:       http.MethodPost,
			URI:          "https://api.mch.weixin.qq.com/pay/unifiedorder",
			ResponseType: "xml",
		},
	},
}
