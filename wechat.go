package wechat

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"wechat/core"
)

// Client 微信接口客户端
type Client struct {
	kernel *core.Kernel
	log    core.Log

	/*
	* certificate key 值如下:
	* 	appid 	开发者ID(AppID)
	*		secret 	开发者密码(AppSecret)
	*   token 	令牌(Token)
	*   aeskey 	消息加解密密钥 (EncodingAESKey)
	 */
	certificate map[string]string
}

// NewClient 初始客户端
func NewClient(certificate map[string]string, log core.Log) *Client {
	cli := Client{
		log:         log,
		kernel:      core.NewKernel(certificate),
		certificate: certificate,
	}

	cli.kernel.SetLogInst(log)
	cli.kernel.StartTokenServer()

	return &cli
}

// GetAppID 获取 AppID
func (t *Client) GetAppID() string {
	return t.certificate["appid"]
}

// Signature 初始校验
func (t *Client) Signature(timestamp, nonce string) string {
	token := t.certificate["token"]
	strs := sort.StringSlice{token, timestamp, nonce}
	sort.Strings(strs)
	str := ""
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

var request core.Request

func init() {
	request = &core.DefaultRequest{}
}
