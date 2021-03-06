package interactor

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type SkuGetter interface {
	GetSkus() ([]domain.Sku, error)
}

type CSVFileSkuGetter struct {
	filename string
}

func NewCSVFileSkuGetter(filename string) *CSVFileSkuGetter {
	return &CSVFileSkuGetter{filename: filename}
}

func (g CSVFileSkuGetter) GetSkus() (skus []domain.Sku, err error) {
	f, err := os.OpenFile(g.filename, os.O_RDONLY, 0777)
	if err != nil {
		return
	} else {
		defer func() { _ = f.Close() }()
	}

	reader := csv.NewReader(f)
	reader.Comma = ';'
	skuStrings, err := reader.ReadAll()
	if err != nil {
		return
	}

	for _, record := range skuStrings {
		if len(record) != 2 {
			return nil, errors.New("В строке должно быть ровно значения, а сейчас " + strconv.Itoa(len(record)))
		}
		skus = append(skus, domain.Sku{
			Id:  record[0],
			Url: record[1],
		})
	}

	return
}

type CSVBytesSkuGetter struct {
	data []byte
}

func NewCSVBytesSkuGetter(data []byte) *CSVBytesSkuGetter {
	return &CSVBytesSkuGetter{data: data}
}

func (g CSVBytesSkuGetter) GetSkus() (skus []domain.Sku, err error) {
	reader := csv.NewReader(bytes.NewReader(g.data))
	reader.Comma = ';'
	skuStrings, err := reader.ReadAll()
	if err != nil {
		return
	}

	for _, record := range skuStrings {
		if len(record) != 2 {
			return nil, errors.New("В строке должно быть ровно значения, а сейчас " + strconv.Itoa(len(record)))
		}
		skus = append(skus, domain.Sku{
			Id:  record[0],
			Url: record[1],
		})
	}

	return
}

func GetHTML(sku domain.Sku, httpClient *http.Client) (body []byte, err error) {
	req := http.Request{
		Method: "GET",
		Host:   "www.wildberries.ru",
	}
	req.URL, err = url.Parse(sku.GetUrl())
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"User-Agent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:88.0) Gecko/20100101 Firefox/88.0"},
		"Accept":     []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},
		"Connection": []string{"keep-alive"},
	}
	resp, err := httpClient.Do(&req)

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
