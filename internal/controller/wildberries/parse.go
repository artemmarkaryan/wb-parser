package wildberries

import (
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/parser"
)

func (w wildberriesController) ParseHTML(
	htmlCh chan *domain.HtmlBody,
	infoCh chan domain.Info,
	errorCh chan error,
) {
	p := parser.WildberriesParser{}

	select {
	case htmlBody := <-htmlCh:
		info, err := p.GetInfo(htmlBody)
		if err != nil {
			errorCh <- err
		} else {
			infoCh <- info
		}
	}
}
