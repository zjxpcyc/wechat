package wechat

import (
	"net/url"
	"wechat/core"
)

// GetOpenID 获取用户 OpenID
func (t *Client) GetOpenID(code string) (string, error) {
	t.log.Info("获取用户 OpenID: code=" + code)

	api := core.WxAPI["oauth2"]["access_token"]
	params := url.Values{}
	params.Set("appid", t.certificate["appid"])
	params.Set("secret", t.certificate["secret"])
	params.Set("code", code)

	res, err := request.GetJSON(api, params)
	if err != nil {
		t.log.Error("获取 Oauth2 Access-Token 失败, ", err.Error())
		return "", err
	}

	return res["openid"].(string), nil
}

// GetUserInfo 获取用户详情
func (t *Client) GetUserInfo(code string) (map[string]interface{}, error) {
	t.log.Info("获取用户详情: code=" + code)

	// 依据 code 获取 openid, access_token
	api := core.WxAPI["oauth2"]["access_token"]
	params := url.Values{}
	params.Set("appid", t.certificate["appid"])
	params.Set("secret", t.certificate["secret"])
	params.Set("code", code)

	res, err := request.GetJSON(api, params)
	if err != nil {
		t.log.Error("获取 Oauth2 Access-Token 失败, ", err.Error())
		return nil, err
	}

	openID := res["openid"].(string)
	token := res["access_token"].(string)

	// 再依据 openid, access_token 获取详情
	api = core.WxAPI["oauth2"]["userinfo"]
	params = url.Values{}
	params.Set("access_token", token)
	params.Set("openid", openID)
	res, err = request.GetJSON(api, params)
	if err != nil {
		t.log.Error("获取 Oauth2 用户信息 失败, ", err.Error())
		return nil, err
	}

	return res, nil
}
