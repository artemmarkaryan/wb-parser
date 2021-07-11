package base_controller

import (
	"bytes"
	"context"
	"github.com/artemmarkaryan/wb-parser/internal/controller/interface"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	makeHTTPClient "github.com/artemmarkaryan/wb-parser/pkg/make-http-client"
	"sync"
)

/*
BaseController — базовый класс, в котором определён только шаблонный метод Process.
Он НЕ реализует интерфейс Controller, и это нормально.
*/
type BaseController struct {
	connectionPoolSize uint8                  // количество потоков для запросов
	parserPoolSize     uint8                  // количество потоков для запросов
	realController     _interface.IController // реализация шаблонного метода
}

func NewBaseController(connectionPoolSize uint8, parserPoolSize uint8, realController _interface.IController) *BaseController {
	return &BaseController{connectionPoolSize: connectionPoolSize, parserPoolSize: parserPoolSize, realController: realController}
}

func (r *BaseController) SetConnectionPoolSize(size uint8) {
	r.connectionPoolSize = size
}

func (r *BaseController) SetParserPoolSize(size uint8) {
	r.parserPoolSize = size
}

func (r *BaseController) getConnectionPoolSize(max uint8) uint8 {
	if max < r.connectionPoolSize {
		return max
	} else {
		return r.connectionPoolSize
	}
}

func (r *BaseController) getParserPoolSize(max uint8) uint8 {
	if max < r.parserPoolSize {
		return max
	} else {
		return r.parserPoolSize
	}
}



// Process описывает логику парсера
// Конкретные функции реализованы в r.realController
func (r BaseController) Process(data *[]byte) (*bytes.Buffer, error) {
	ctx := context.Background()

	skus, err := r.realController.ConvertBytesToSku(data)
	if err != nil {
		return nil, err
	}

	skuCh := make(chan domain.Sku)
	go func() {
		for _, s := range *skus {
			skuCh <- s
		}
		close(skuCh)
	}() // put sku in channel

	infoCh := make(chan *domain.Info)
	htmlCh := make(chan *domain.HtmlBody)
	buffCh := make(chan *bytes.Buffer)

	// Requesting
	clientCtx := context.WithValue(ctx, "client", makeHTTPClient.NewHTTPClient(len(*skus)))
	requestWG := sync.WaitGroup{}
	requestErrCh := make(chan error, len(*skus))
	for i := uint8(0); i < r.getConnectionPoolSize(uint8(len(*skus))); i++ {
		requestWG.Add(1)
		go r.realController.Request(
			clientCtx,
			skuCh,
			htmlCh, // используется в Parsing
			requestErrCh,
			&requestWG,
		)
	}

	// Parsing
	parseErrCh := make(chan error, len(*skus))
	parseWG := sync.WaitGroup{}
	for i := uint8(0); i < r.getParserPoolSize(uint8(len(*skus))); i++ {
		parseWG.Add(1)
		go r.realController.ParseHTML(
			htmlCh,
			infoCh, // используется в Exporting
			parseErrCh,
			&parseWG,
		)
	}

	// Exporting
	exportErrCh := make(chan error, 1)
	go r.realController.Export(ctx, infoCh, buffCh, exportErrCh)

	// Sync
	go func() {
		requestWG.Wait()
		close(htmlCh)
	}()

	go func() {
		parseWG.Wait()
		close(infoCh)
	}()

	select {
	case buff := <-buffCh:
		return buff, nil
	}
}
