package ozon

import (
	"bytes"
	"context"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/exporter"
)

func (c ozonController) Export(
	ctx context.Context,
	infoCh chan *domain.Info,
	buffCh chan *bytes.Buffer,
	errorCh chan error,
) {
	var processedData []byte
	buff := bytes.NewBuffer(processedData)
	err := exporter.ExcelExporter{}.Export(ctx, infoCh, buff)
	buffCh <- buff
	errorCh <- err
}
