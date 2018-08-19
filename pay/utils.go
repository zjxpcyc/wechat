package pay

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Struct2Map struct 转 map
func Struct2Map(s interface{}) (map[string]interface{}, error) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	// 如果是指针
	if t.Kind() == reflect.Ptr {
		t = v.Type()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, errors.New("Transform struct to map fail: not a struct")
	}

	res := make(map[string]interface{})

	// 遍历字段
	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		ft := t.Field(i)
		tag := ft.Tag.Get("xml")

		key := ft.Name
		tagDetail := strings.Split(tag, ",")
		if len(tagDetail) > 0 {
			key = tagDetail[0]
		}

		val := fv.Interface()

		// 如果有需要忽略的字段
		if strings.Index(tag, "omitempty") > -1 {
			// 微信支付接口中, 基础类型只有 string 与 int
			switch directV := val.(type) {
			case string:
				if directV == "" {
					continue
				}
			case int:
				if directV == 0 {
					continue
				}
			default:
			}
		}

		// 复杂的类型, 转成 json
		if strings.Index(tag, "tojson") > -1 {
			bf := bytes.NewBuffer([]byte{})
			encode := json.NewEncoder(bf)
			if err := encode.Encode(val); err != nil {
				return nil, err
			}

			valStr := bf.String()
			if strings.Index(tag, "withcdata") > -1 {
				valStr = fmt.Sprintf("<![CDATA[%s]]>", valStr)
			}

			res[key] = valStr
			continue
		}

		// 如果需要 cdata 包裹
		if fv.Kind() == reflect.String && strings.Index(tag, "withcdata") > -1 {
			res[key] = fmt.Sprintf("<![CDATA[%v]]>", val)
			continue
		}

		res[key] = val
	}

	return res, nil
}
