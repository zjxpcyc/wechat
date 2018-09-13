package wx

import (
	"net/url"
	"time"
)

// JsTicketTask 刷新 任务
func (t *Client) JsTicketTask() time.Duration {
	var reTrySec int64 = 60
	ticket, expire, err := t.getJsTicket()
	if err != nil {
		log.Error("获取 JS Ticket 失败", err.Error())
		expire = reTrySec
	}

	// 不允许连续不断调用
	if expire == 0 {
		expire = reTrySec
	}

	t.jsTicket = ticket
	return time.Duration(expire) * time.Second
}

// getJsTicket 获取 token
func (t *Client) getJsTicket() (string, int64, error) {
	api := API["jssdk"]["ticket"]
	params := url.Values{}
	params.Set("access_token", t.accessToken)

	res, err := t.request.Do(api, params)
	if err != nil {
		return "", 0, err
	}

	ticket, _ := res["ticket"].(string)
	expire, _ := res["expires_in"].(float64)
	return ticket, int64(expire), nil
}

func (t *Client) JsTicketSign(url string) (map[string]string, error) {
	return nil, nil
}
