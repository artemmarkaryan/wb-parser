package ozon

import (
	"context"
	"errors"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/html_retriever"
	"log"
	"net/http"
	"sync"
)

var NoClientInContext = errors.New("no client in context")

func (c ozonController) Request(
	ctx context.Context,
	skuCh chan domain.Sku,        // read
	htmlCh chan *domain.HtmlBody, // write
	errCh chan error,             // write
	syncMap *sync.Map,
) {
	var clientErr error
	retriever := html_retriever.OzonHtmlRetriever{}

	client, clientOk := ctx.Value("client").(*http.Client)
	if !clientOk {
		clientErr = NoClientInContext
		log.Panic(clientErr)
	}

	select {
	case sku, open := <-skuCh:
		if !open {
			break
		}
		htmlBodyBytes, err := retriever.GetHTML(sku, client)
		log.Printf("recieved from %v", sku.GetId())
		htmlBody := domain.HtmlBody(htmlBodyBytes)
		if err != nil {
			errCh <- err
		} else {
			htmlCh <- &htmlBody
		}
	}
}
