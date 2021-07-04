package exporter

import (
	"context"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"io"
)

type IExporter interface {
	Export(context.Context, chan *domain.Info, io.Writer) error
}
