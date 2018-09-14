package core

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

// URLValues2XMLString url.Values 转 xml 字串
func URLValues2XMLString(dt url.Values) string {
	var res string
	for k := range dt {
		res = res + fmt.Sprintf("<%s>%s</%s>", k, dt.Get(k), k)
	}

	return "<xml>" + res + "</xml>"
}

// RandomIntn 随机数
func RandomIntn(length int, max int) []int {
	seed := rand.NewSource(int64(time.Now().Nanosecond() + rand.Intn(6)))
	r := rand.New(seed)

	res := make([]int, 0)
	for i := 0; i < length; i++ {
		res = append(res, r.Intn(max))
	}

	return res
}
