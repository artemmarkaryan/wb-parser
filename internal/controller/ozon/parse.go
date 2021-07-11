package ozon

import (
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"sync"
)

func (c ozonController) ParseHTML(
	htmlCh chan *domain.HtmlBody, // read
	infoCh chan *domain.Info,     // write
	errorCh chan error,           // write
	wg *sync.WaitGroup,
) {
loop:
	for {
		select {
		case html, open := <-htmlCh:
			if !open {
				break loop
			}
			infoCh <- &domain.Info{"html": string(*html)}
		}
	}

	wg.Done()
}
