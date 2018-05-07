package wechat

import (
	"net/url"
	"time"
)

// Task Access-Token 刷新 任务
func (t *Client) Task() time.Duration {
	var reTrySec int64 = 60
	token, expire, err := t.getToken()
	if err != nil {
		log.Error("获取 Access-Token 失败", err.Error())
		expire = reTrySec
	}

	t.accessToken = token
	return time.Duration(expire)
}

// getToken 获取 token
func (t *Client) getToken() (string, int64, error) {
	api := WxAPI["access_token"]["get"]
	params := url.Values{}
	params.Set("appid", t.certificate["appid"])
	params.Set("secret", t.certificate["secret"])

	res, err := t.request.GetJSON(api, params)
	if err != nil {
		return "", 0, err
	}

	token := res["access_token"].(string)
	expire := res["expires_in"].(float64)
	return token, int64(expire), nil
}
