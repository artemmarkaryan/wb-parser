package parser

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetInfo(t *testing.T) {
	resp, _ := http.Get("https://www.wildberries.ru/catalog/6031342/detail.aspx")
	t.Logf(resp.Status)
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	t.Logf("bytes: %v", string(bodyBytes[:100]))
	infos, err := WildberriesParser{}.GetInfo(bodyBytes)

	for title, value := range infos {
		t.Logf("{%q: %q}", title, value)
	}
	if err != nil {
		t.Error(err.Error())
	}
}
