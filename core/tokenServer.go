package core

import (
	"net/url"
	"time"
)

// AccessTokenServer  access_token 获取服务
type AccessTokenServer interface {
	// Start 启动服务
	Start()

	// Stop 停止服务
	Stop()

	// Get 获取 access_token
	Get() string
}

// DefaultTokenServer 默认 access_token 获取服务
type DefaultTokenServer struct {
	done        chan bool
	token       string
	certificate map[string]string
	request     *DefaultRequest
}

// NewDefaultTokenServer 构造默认 Token Server
func NewDefaultTokenServer(certificate map[string]string) *DefaultTokenServer {
	return &DefaultTokenServer{
		done:        make(chan bool),
		token:       "",
		certificate: certificate,
		request:     &DefaultRequest{},
	}
}

// Get 获取 access_token
func (t *DefaultTokenServer) Get() string {
	return t.token
}

// Stop 停止服务
func (t *DefaultTokenServer) Stop() {
	t.done <- true
}

// Start 启动服务
func (t *DefaultTokenServer) Start() {
	go t.task()

	for {
		select {
		case done := <-t.done:
			if done {
				return
			}
		}
	}
}

// task 定时任务
func (t *DefaultTokenServer) task() {
	// 发生错误 60 秒后重试
	var reTry int64 = 60
	token, expire, err := t.getToken()
	if err != nil {
		expire = reTry
	}

	t.token = token
	log.Info("TokenServer 获取到 access_token: " + token)
	ticker := time.NewTicker(time.Second * time.Duration(expire))

	for {
		select {
		case <-ticker.C:
			ticker.Stop()

			token, expire, err = t.getToken()
			if err != nil {
				expire = reTry
			}

			t.token = token
			log.Info("TokenServer 获取到 access_token: " + token)
			ticker = time.NewTicker(time.Second * time.Duration(expire))
		}
	}
}

// getToken 获取 token
func (t *DefaultTokenServer) getToken() (string, int64, error) {
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

var _ AccessTokenServer = &DefaultTokenServer{}
