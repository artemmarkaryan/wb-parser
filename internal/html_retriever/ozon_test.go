package html_retriever

import (
	"fmt"
	"github.com/andybalholm/brotli"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestRequest(t *testing.T) {
	req := &http.Request{
		Method: "GET",
		Host:   "www.ozon.ru",
	}
	req.URL = &url.URL{
		Scheme: "https",
		Host:   "www.ozon.ru",
		Path:   "/context/detail/id/166562057",
	}
	req.Header = http.Header{
		"User-Agent":      []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:89.0) Gecko/20100101 Firefox/89.0"},
		"Accept":          []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},
		"Accept-Encoding": []string{"gzip,deflate,br"},
		"Connection":      []string{"keep-alive"},
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err.Error())
		return
	}

	//bodyBytes, err := ioutil.ReadAll(resp.Body)
	//t.Logf("%v", bodyBytes)
	r0 := brotli.NewReader(resp.Body)
	r1, _ := ioutil.ReadAll(r0)
	_ = ioutil.WriteFile(
		fmt.Sprintf("./temp%v.html", time.Now().Unix()),
		r1, 0644)
	//t.Log(string(r1))

	//log.Print(string(res))
}
