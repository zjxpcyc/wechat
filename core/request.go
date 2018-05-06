package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Request 可以直接使用的 http request
type Request interface {
	// GetJSON GET 远程数据, 并返回 json
	GetJSON(APIInfo, url.Values) (map[string]interface{}, error)
}

// DefaultRequest 简易 http request
type DefaultRequest struct{}

// GetJSON GET 远程数据, 并返回 json
func (t *DefaultRequest) GetJSON(api APIInfo, params url.Values) (map[string]interface{}, error) {
	apiURL, _ := url.Parse(APIDomain + api.URI)
	query := apiURL.Query()

	for k := range params {
		query.Set(k, params.Get(k))
	}

	apiURL.RawQuery = query.Encode()
	remoteAddr := apiURL.String()

	log.Info("请求远程接口: ", remoteAddr)

	resp, err := http.Get(remoteAddr)
	if err != nil {
		log.Error("请求远程数据失败 (GET: "+remoteAddr+")", err.Error())
		return nil, err
	}

	var res map[string]interface{}
	res, err = t.jsonResult(resp)
	if err != nil {
		return nil, err
	}

	return res, t.checkJSONResult(res)
}

func (t *DefaultRequest) jsonResult(resp *http.Response) (map[string]interface{}, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("解析请求结果失败, ", err.Error())
		return nil, err
	}

	var res map[string]interface{}
	if err := json.Unmarshal(body, &res); err != nil {
		log.Error("转换请求结果(JSON)失败", err.Error())
		return nil, err
	}

	return res, nil
}

func (t *DefaultRequest) checkJSONResult(res map[string]interface{}) error {
	log.Info("接口返回结果: ", res)

	errcode, _ := res["errcode"]
	errmsg, _ := res["errmsg"]
	if errcode == nil {
		return nil
	}

	err, _ := errcode.(float64)
	errNum := int(err)

	if errNum == 0 {
		return nil
	}

	msg, _ := errmsg.(string)
	log.Error("接口返回错误: " + strconv.Itoa(errNum) + "-" + msg)
	return errors.New(msg)
}

var _ Request = &DefaultRequest{}
