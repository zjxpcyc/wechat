package core

import (
	"time"
)

// ScheduleTask 定时任务
type ScheduleTask interface {
	// Start 启动服务
	Start()

	// Stop 停止服务
	Stop()
}

// Task 任务, 返回下次执行的时间间隔
type Task func() time.Duration

// DefaultScheduleServer 默认 access_token 获取服务
type DefaultScheduleServer struct {
	done chan bool
	task Task
}

// NewDefaultScheduleServer 构造默认 Token Server
func NewDefaultScheduleServer() *DefaultScheduleServer {
	return &DefaultScheduleServer{
		done: make(chan bool),
	}
}

// SetTask 设置定时任务
func (t *DefaultScheduleServer) SetTask(task Task) {
	t.task = task
}

// Stop 停止服务
func (t *DefaultScheduleServer) Stop() {
	t.done <- true
}

// Start 启动服务
func (t *DefaultScheduleServer) Start() {
	go func() {
		d := t.task()
		for {
			time.Sleep(d)
			d = t.task()
		}
	}()

	for {
		select {
		case done := <-t.done:
			if done {
				return
			}
		}
	}
}

// // task 定时任务
// func (t *DefaultTokenServer) task() {
// 	// 发生错误 60 秒后重试
// 	var reTry int64 = 60
// 	token, expire, err := t.getToken()
// 	if err != nil {
// 		expire = reTry
// 	}

// 	t.token = token
// 	log.Info("TokenServer 获取到 access_token: " + token)
// 	ticker := time.NewTicker(time.Second * time.Duration(expire))

// 	for {
// 		select {
// 		case <-ticker.C:
// 			ticker.Stop()

// 			token, expire, err = t.getToken()
// 			if err != nil {
// 				expire = reTry
// 			}

// 			t.token = token
// 			log.Info("TokenServer 获取到 access_token: " + token)
// 			ticker = time.NewTicker(time.Second * time.Duration(expire))
// 		}
// 	}
// }

// // getToken 获取 token
// func (t *DefaultTokenServer) getToken() (string, int64, error) {
// 	api := WxAPI["access_token"]["get"]
// 	params := url.Values{}
// 	params.Set("appid", t.certificate["appid"])
// 	params.Set("secret", t.certificate["secret"])

// 	res, err := t.request.GetJSON(api, params)
// 	if err != nil {
// 		return "", 0, err
// 	}

// 	token := res["access_token"].(string)
// 	expire := res["expires_in"].(float64)
// 	return token, int64(expire), nil
// }

var _ ScheduleTask = &DefaultScheduleServer{}
