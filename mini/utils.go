package mini

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/lunny/log"
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

// Decrypt 解密小程序返回数据
// https://developers.weixin.qq.com/miniprogram/dev/api/signature.html#wxchecksessionobject
func Decrypt(dt, iv, key string) (map[string]interface{}, error) {
	// 转换 iv
	rawIv, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	aesIV := []byte(rawIv)

	// 转换 session_key
	rawKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	aesKey := []byte(rawKey)

	// 转换加密数据
	data, err := base64.StdEncoding.DecodeString(dt)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, aesIV)

	dist := make([]byte, len(data))
	mode.CryptBlocks(dist, data)
	// dist = pKCS7UnPadding(dt)

	var res map[string]interface{}
	if err := json.Unmarshal(dist, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func pKCS7UnPadding(dt []byte) []byte {
	length := len(dt)
	unpadding := int(dt[length-1])
	return dt[:(length - unpadding)]
}
