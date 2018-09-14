package mini

import (
	"github.com/zjxpcyc/wechat/core"
)

var log core.Log

// Client 微信小程序接口客户端
type Client struct {
	kernel  *core.Kernel
	request core.Request

	/*
	* certificate key 值如下:
	* 	appid 	开发者ID(AppID)
	*		secret 	开发者密码(AppSecret)
	 */
	certificate map[string]string
}

// NewClient 初始客户端
func NewClient(certificate map[string]string, log core.Log) *Client {
	cli := &Client{
		request:     core.NewDefaultRequest(checkJSONResult),
		kernel:      core.NewKernel(),
		certificate: certificate,
	}

	return cli
}

// SetLogInst 设置全局日志实例
func SetLogInst(l core.Log) {
	core.SetLogInst(l)
	log = l
}

func init() {
	log = &core.DefaultLogger{}
}
