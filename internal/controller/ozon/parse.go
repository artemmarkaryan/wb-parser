package ozon

import (
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"log"
	"sync"
)

func (c ozonController) ParseHTML(
	htmlCh chan *domain.HtmlBody, // read
	infoCh chan *domain.Info,     // write
	errorCh chan error,           // write
	syncMap *sync.Map,
) {
	select {
	case html, open := <-htmlCh:
		if !open {
			syncMap.
			break
		}
		log.Print((*html)[:1])
		infoCh <- &domain.Info{}
	}
}
