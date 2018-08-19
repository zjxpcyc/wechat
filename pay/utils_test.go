package pay_test

import (
	"fmt"
	"testing"

	"github.com/zjxpcyc/wechat/pay"
)

func TestStruct2Map(t *testing.T) {
	testVal := pay.ParamsUnifiedOrder{
		Body: "JSAPI支付测试",
		Detail: pay.PreferentialDetail{
			GoodsDetail: []pay.GoodsDetail{
				pay.GoodsDetail{
					GoodsID:       "iphone6s_16G",
					WxpayGoodsID:  "1001",
					GoodsName:     "iPhone6s 16G",
					Quantity:      1,
					Price:         528800,
					GoodsCategory: "123456",
					Body:          "苹果手机",
				},
				pay.GoodsDetail{
					GoodsID:       "iphone6s_32G",
					WxpayGoodsID:  "1002",
					GoodsName:     "iPhone6s 32G",
					Quantity:      1,
					Price:         608800,
					GoodsCategory: "123789",
					Body:          "苹果手机",
				},
			},
		},
		NotifyURL:      "http://wxpay.wxutil.com/pub_v2/pay/notify.v2.php",
		OpenID:         "oUpF8uMuAJO_M2pxb1Q9zNjWeS6o",
		OutTradeNo:     "1415659990",
		SpbillCreateIP: "14.23.150.211",
		TotalFee:       1,
		TradeType:      "JSAPI",
	}

	// expected := ""

	xmlstr, err := pay.Struct2Map(testVal)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	fmt.Println(xmlstr)

	// if xmlstr != expected {
	// 	t.Fatalf("Test Struct2Map func fail")
	// }
}
