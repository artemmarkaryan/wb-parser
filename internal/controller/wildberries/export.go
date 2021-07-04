package wildberries

import (
	"bytes"
	"context"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/exporter"
)

func (w wildberriesController) Export(ctx context.Context, infoCh chan *domain.Info) (*bytes.Buffer, error) {
	var processedData []byte
	buff := bytes.NewBuffer(processedData)
	err := exporter.ExcelExporter{}.Export(ctx, infoCh, buff)

	return buff, err
}

