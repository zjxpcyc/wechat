package mini

import (
	"net/url"

	"github.com/zjxpcyc/wechat/core"
)

// GetOpenID 获取用户 OpenID
func (t *Client) GetOpenID(code string) (map[string]interface{}, error) {
	log.Info("获取用户 OpenID: code=" + code)

	api := API["oauth2"]["session"]
	params := url.Values{}
	params.Set("appid", t.certificate["appid"])
	params.Set("secret", t.certificate["secret"])
	params.Set("js_code", code)

	res, err := t.request.Do(api, params)
	if err != nil {
		log.Error("获取 登录凭证 失败, ", err.Error())
		return nil, err
	}

	return res, nil
}

// GetUserFromEncryptData 解析加密数据
func (t *Client) GetUserFromEncryptData(encryptedData, sessionKey, iv string) (map[string]interface{}, error) {
	res, err := core.DecodeMiniData(encryptedData, iv, sessionKey)
	if err != nil {
		log.Error("解密小程序数据失败", err.Error())
		return nil, err
	}

	return res, nil
}
