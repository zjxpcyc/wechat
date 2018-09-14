package pay

import (
	"bytes"
	"errors"

	"github.com/zjxpcyc/wechat/core"
)

// Query 订单查询
func (t *Client) Query(transID string, tradeNO ...string) (string, map[string]interface{}, error) {
	xmlBody, err := t.prepareQueryParams(transID, tradeNO...)
	if err != nil {
		return "", nil, err
	}

	api := API["order"]["query"]
	data := bytes.NewBuffer([]byte(xmlBody))

	// 查询请求
	resp, err := t.request.Do(api, nil, data)
	if err != nil {
		return "", nil, err
	}

	log.Info("支付订单查询结果: ", resp)

	// 通讯
	rtnCode := resp["return_code"].(string)
	if rtnCode != ResponseSuccess {
		msg := resp["return_msg"].(string)
		return "", nil, errors.New(msg)
	}

	// 业务
	resCode := resp["result_code"].(string)
	if resCode != BizSuccess {
		msg := resp["err_code_des"].(string)
		return "", nil, errors.New(msg)
	}

	// 交易
	tradeState := resp["trade_state"].(string)
	if tradeState == TradePayError {
		msg := resp["trade_state_desc"].(string)
		return "", nil, errors.New(msg)
	}

	return tradeState, resp, nil
}

// tradeNO, transID 只能是二选一
func (t *Client) prepareQueryParams(transID string, tradeNO ...string) (string, error) {
	outTradeNO := ""
	if len(tradeNO) > 0 && tradeNO[0] != "" {
		outTradeNO = tradeNO[0]
	}

	if transID != "" {
		outTradeNO = ""
	}

	params := map[string]interface{}{
		"appid":          t.certificate["appid"],
		"mch_id":         t.certificate["mch_id"],
		"transaction_id": transID,
		"out_trade_no":   outTradeNO,
	}

	signMap := core.GetSignOfPay(params)
	params["sign"] = signMap["sign"]
	params["nonce_str"] = signMap["nonce_str"]

	return core.Map2XMLString(params), nil
}
