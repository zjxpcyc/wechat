package core

import (
	"fmt"
	"net/url"
)

// URLValues2XMLString url.Values 转 xml 字串
func URLValues2XMLString(dt url.Values) string {
	var res string
	for k := range dt {
		res = res + fmt.Sprintf("<%s>%s</%s>", k, dt.Get(k), k)
	}

	return "<xml>" + res + "</xml>"
}
