package wx

import (
	"errors"
	"strconv"
)

func checkJSONResult(res map[string]interface{}) error {
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
