// interactor is an interface to database
package interactor

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// get sku instances from some source
type SkuGetter interface {
	GetSkus() ([]domain.Sku, error)
}

type CSVSkuGetter struct {
	filename string
}

func NewCSVSkuGetter(filename string) *CSVSkuGetter {
	return &CSVSkuGetter{filename: filename}
}

func (g CSVSkuGetter) GetSkus() (skus []domain.Sku, err error) {
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

// todo: make excel sku getter

// retrieve html data
func GetHTML(sku domain.Sku, httpClient *http.Client) (body io.Reader, err error) {
	log.Print("requesting ", sku.GetUrl())

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
		"Accept": []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},
		"Connection": []string{"keep-alive"},
	}
	resp, err := httpClient.Do(&req)

	if err != nil {
		return
	}
	//defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		err = errors.New(
			fmt.Sprintf("response to %v resuled in %v", sku.Url, resp.StatusCode),
		)
	}

	body = resp.Body
	return
}
