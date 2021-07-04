package html_retriever

import (
	"errors"
	"fmt"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"io/ioutil"
	"net/http"
	"net/url"
)

type OzonHtmlRetriever struct{}

func (r OzonHtmlRetriever) configureRequest(sku domain.Sku) (req *http.Request, err error) {
	req = &http.Request{
		Method: "GET",
		Host:   "www.ozon.ru",
	}
	//req.URL = &url.URL{
	//	Scheme: "https",
	//	Host:   "www.ozon.ru",
	//	Path:   "/context/detail/id/166562057",
	//}
	req.URL, err = url.Parse(sku.GetUrl())
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"User-Agent":      []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:89.0) Gecko/20100101 Firefox/89.0"},
		"Accept":          []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},
		"Accept-Encoding": []string{"gzip,deflate,br"},
		"Connection":      []string{"keep-alive"},
	}
	return
}

func (r OzonHtmlRetriever) GetHTML(sku domain.Sku, httpClient *http.Client) (body []byte, err error) {
	req, err := r.configureRequest(sku)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = resp.Body.Close()
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(
			fmt.Sprintf("response to %v resuled in %v", sku.Url, resp.StatusCode),
		)
	}

	return bodyBytes, err
}
