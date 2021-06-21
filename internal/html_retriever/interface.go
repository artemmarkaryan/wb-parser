package html_retriever

import (
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"net/http"
)

type HtmlRetriever interface {
	GetHTML(sku domain.Sku, httpClient *http.Client) (body []byte, err error)
}