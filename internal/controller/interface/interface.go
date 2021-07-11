// Контроллер контролирует процесс парсинга

package _interface

import (
	"bytes"
	"context"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"sync"
)

type IController interface {
	ConvertBytesToSku(*[]byte) (*[]domain.Sku, error)
	Request(
		ctx context.Context,
		sku chan domain.Sku,        // read
		html chan *domain.HtmlBody, // write
		err chan error,             // write
		wg *sync.WaitGroup,
	)
	ParseHTML(
		html chan *domain.HtmlBody, // read
		info chan *domain.Info,     // write
		err chan error,             // write
		wg *sync.WaitGroup,
	)
	Export(
		ctx context.Context,
		info chan *domain.Info,  // read
		buff chan *bytes.Buffer, // write
		err chan error,          // write
	)
	// Process — шаблонный метод
	Process(*[]byte) (*bytes.Buffer, error)
}
