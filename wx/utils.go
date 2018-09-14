package wx

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/zjxpcyc/wechat/core"
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
	return errors.New(strconv.Itoa(errNum) + "-" + msg)
}

// RandomString 随机字符
func RandomString(l int) string {
	serial := strings.Split("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

	nums := core.RandomIntn(l, len(serial))

	res := []string{}
	for i := 0; i < l; i++ {
		res = append(res, serial[nums[i]])
	}

	return strings.Join(res, "")
}

// JsTicketSignature 计算 js-ticket signature
func JsTicketSignature(url, noncestr, ticket, timestamp string) string {
	willSign := []string{
		"noncestr=" + noncestr,
		"timestamp=" + timestamp,
		"url=" + url,
		"jsapi_ticket=" + ticket,
	}
	sort.Strings(willSign)
	str2Sign := strings.Join(willSign, "&")

	return Sha1(str2Sign)
}

//Sha1  sha1 小写
func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
