package core_test

import (
	"net/url"
	"testing"

	"github.com/zjxpcyc/wechat/core"
)

func TestURLValues2XMLString(t *testing.T) {
	query := "a=b&c=d"
	expected := []string{
		"<xml><a>b</a><c>d</c></xml>",
		"<xml><c>d</c><a>b</a></xml>",
	}
	testData, _ := url.ParseQuery(query)

	res := core.URLValues2XMLString(testData)

	if res != expected[0] && res != expected[1] {
		t.Fatalf("Transfrom url.Values to xml string fail, %s", res)
	}
}
