// interactor is an interface to database
package interactor

import (
	"errors"
	"fmt"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"io"
	"log"
	"net/http"
)

func GetHTML(sku domain.Sku, httpClient *http.Client) (body io.ReadCloser, err error) {
	log.Print("requesting ", sku.GetUrl())
	resp, err := httpClient.Get(sku.GetUrl())
	if err != nil {
		return
	}
	//defer func() {_ = resp.Body.Close()}()

	if resp.StatusCode != 200 {
		err = errors.New(
			fmt.Sprintf("response to %v resuled in %v", sku.Url, resp.StatusCode),
		)
	}
	return resp.Body, err
}