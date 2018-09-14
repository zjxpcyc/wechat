package pay

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

// GetPayResult 支付完成回调
func (t *Client) GetPayResult(r *http.Request) (map[string]interface{}, error) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	resp, err := t.getRequestBody(r)
	if err != nil {
		return nil, err
	}

	log.Info("支付通知结果: ", resp)

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

// getRequestBody 获取事件结果
func (t *Client) getRequestBody(r *http.Request) (map[string]interface{}, error) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var rtn map[string]interface{}
	if err := xml.Unmarshal(resp, &rtn); err != nil {
		return nil, err
	}

	return rtn, nil
}
