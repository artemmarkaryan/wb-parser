package parser

import (
	"github.com/artemmarkaryan/wb-parser/internal/config"
	"net/http"
	"testing"
)

func TestGetInfo(t *testing.T) {

	_ = config.LoadDotEnv()
	resp, _ := http.Get("https://www.wildberries.ru/catalog/6031342/detail.aspx")

	infos, err := GetInfo(resp.Body)

	for title, value := range infos {
		t.Logf("{%q: %q}", title, value)
	}
	if err != nil {
		t.Error(err.Error())
	}
}
