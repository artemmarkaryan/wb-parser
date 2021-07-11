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
	wg *sync.WaitGroup,
) {
	var clientErr error
	retriever := html_retriever.OzonHtmlRetriever{}

	client, clientOk := ctx.Value("client").(*http.Client)
	if !clientOk {
		clientErr = NoClientInContext
		log.Panic(clientErr)
	}

loop:
	for {
		select {
		case sku, open := <-skuCh:
			if !open {
				break loop
			}
			htmlBodyBytes, err := retriever.GetHTML(sku, client)
			htmlBody := domain.HtmlBody(htmlBodyBytes)
			if err != nil {
				errCh <- err
			} else {
				htmlCh <- &htmlBody
			}

		}
	}

	wg.Done()
}
