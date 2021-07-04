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

	syncMap := sync.Map{}

	// Requesting
	clientCtx := context.WithValue(
		ctx,
		"client",
		makeHTTPClient.NewHTTPClient(len(*skus)),
	)
	requestErrCh := make(chan error, len(*skus))
	for i := uint8(0); i < r.connectionPoolSize; i++ {
		go r.realController.Request(
			clientCtx,
			skuCh,
			htmlCh, // используется в Parsing
			requestErrCh,
			&syncMap,
		)
	}

	// Parsing
	parseErrCh := make(chan error, len(*skus))
	for i := uint8(0); i < r.parserPoolSize; i++ {
		go r.realController.ParseHTML(
			htmlCh,
			infoCh, // используется в Exporting
			parseErrCh,
			&syncMap,
		)
	}

	// Exporting
	exportErrCh := make(chan error, 1)
	go r.realController.Export(ctx, infoCh, buffCh, exportErrCh)

	select {
	case buff := <-buffCh:
		return buff, nil
	}
}
