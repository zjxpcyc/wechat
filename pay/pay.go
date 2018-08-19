package pay

import (
	"github.com/zjxpcyc/wechat/core"
)

// 通讯结果
const (
	ResponseSuccess = "SUCCESS"
	ResponseFail    = "FAIL"
)

// 业务结果
const (
	BizSuccess = "SUCCESS"
	BizFail    = "FAIL"
)

// 交易结果
const (
	TradeSuccess    = "SUCCESS"
	TradeRefund     = "REFUND"
	TradeNotPay     = "NOTPAY"
	TradeClosed     = "CLOSED"
	TradeRevoked    = "REVOKED"
	TradeUserPaying = "USERPAYING"
	TradePayError   = "PAYERROR"
)

var log core.Log

// Client 微信支付接口客户端
type Client struct {
	kernel  *core.Kernel
	request core.Request

	/*
	* certificate key 值如下:
	* 	appid 			公众账号ID(AppID)
	*   mch_id  		商户号
	*		notify_url 	通知地址
	*		key					API安全 密钥
	 */
	certificate map[string]string
}

// GoodsDetail 单品信息
type GoodsDetail struct {
	GoodsID       string `json:"goods_id"`
	WxpayGoodsID  string `json:"wxpay_goods_id,omitempty"`
	GoodsName     string `json:"goods_name,omitempty"`
	GoodsCategory string `json:"goods_category,omitempty"`
	Body          string `json:"body,omitempty"`
	Quantity      int    `json:"quantity"`
	Price         int    `json:"price"`
}

// PreferentialDetail 单品优惠
type PreferentialDetail struct {
	CostPrice   int           `json:"cost_price,omitempty"`
	ReceiptID   string        `json:"receipt_id,omitempty"`
	GoodsDetail []GoodsDetail `json:"goods_detail"`
}

// ParamsUnifiedOrder 下单参数
// 该参数实际上会在 接口方法中被转换为 map
// 此处使用了官方的 xml tag 定义, 但是也增加了部分自定的内容
// withcdata 	需要使用 cdata
// tojson 		字段需要转换为 json
type ParamsUnifiedOrder struct {
	Body           string             `xml:"body"`
	Detail         PreferentialDetail `xml:"detail,omitempty,withcdata,tojson"`
	Attach         string             `xml:"attach,omitempty"`
	OutTradeNo     string             `xml:"out_trade_no"`
	FeeType        string             `xml:"fee_type,omitempty"`
	TotalFee       int                `xml:"total_fee"`
	SpbillCreateIP string             `xml:"spbill_create_ip"`
	TimeStart      string             `xml:"time_start,omitempty"`
	TimeExpire     string             `xml:"time_expire,omitempty"`
	GoodsTag       string             `xml:"goods_tag,omitempty"`
	NotifyURL      string             `xml:"notify_url"`
	TradeType      string             `xml:"trade_type"`
	ProductID      string             `xml:"product_id,omitempty"`
	LimitPay       string             `xml:"limit_pay,omitempty"`
	OpenID         string             `xml:"openid,omitempty"`
	SceneInfo      string             `xml:"scene_info,omitempty,withcdata"`
}

// NewClient 初始客户端
func NewClient(certificate map[string]string) *Client {
	cli := &Client{
		request:     core.NewDefaultRequest(),
		kernel:      core.NewKernel(),
		certificate: certificate,
	}

	return cli
}
