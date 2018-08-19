package pay

import (
	"bytes"
	"errors"

	"github.com/zjxpcyc/wechat/core"
)

// Order 下单
func (t *Client) Order(orderInfo *ParamsUnifiedOrder) (map[string]interface{}, error) {
	xmlBody, err := t.prepareOrderParams(orderInfo)
	if err != nil {
		return nil, err
	}

	api := API["order"]["create"]
	data := bytes.NewBuffer([]byte(xmlBody))

	// 下单请求
	resp, err := t.request.Do(api, nil, data)
	if err != nil {
		return nil, err
	}

	log.Info("支付下单结果: ", resp)

	// 通讯
	rtnCode := resp["return_code"].(string)
	if rtnCode != ResponseSuccess {
		msg := resp["return_msg"].(string)
		return nil, errors.New(msg)
	}

	// 业务
	resCode := resp["result_code"].(string)
	if resCode != BizSuccess {
		msg := resp["err_code_des"].(string)
		return nil, errors.New(msg)
	}

	return resp, nil
}

func (t *Client) prepareOrderParams(orderInfo *ParamsUnifiedOrder) (string, error) {
	// 下单内容转换为 map 类型
	orderMap, err := Struct2Map(orderInfo)
	if err != nil {
		return "", err
	}

	// 添加必要的字段
	orderMap["appid"] = t.certificate["appid"]
	orderMap["mch_id"] = t.certificate["mch_id"]
	if _, ok := orderMap["notify_url"]; !ok {
		orderMap["notify_url"] = t.certificate["notify_url"]
	}
	if _, ok := orderMap["trade_type"]; !ok {
		orderMap["trade_type"] = "JSAPI"
	}

	// 获取 sign
	signMap := core.GetSignOfPay(orderMap)
	orderMap["sign"] = signMap["sign"]
	orderMap["nonce_str"] = signMap["nonce_str"]

	return core.Map2XMLString(orderMap), nil
}
