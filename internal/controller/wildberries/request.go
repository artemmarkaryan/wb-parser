package wildberries

import (
	"context"
	"errors"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/html_retriever"
	"net/http"
)

var NoClientInContext = errors.New("no client in context")

func (w wildberriesController) Request(
	ctx context.Context,
	skuCh chan domain.Sku,        // read from
	htmlCh chan *domain.HtmlBody, // write to
	errCh chan error,             // write to
) {
	var clientErr error
	retriever := html_retriever.WildberriesHtmlRetriever{}

	client, clientOk := ctx.Value("client").(*http.Client)
	if !clientOk {
		clientErr = NoClientInContext
	}


	select {
	case sku, open := <-skuCh:
		if !open {
			close(htmlCh)
			break
		}

		if client == nil {
			errCh <- clientErr
		} else {
			htmlBody, err := retriever.GetHTML(sku, client)
			if err != nil {
				errCh <- err
			} else {
				htmlCh <- htmlBody
			}
		}

	}
}
