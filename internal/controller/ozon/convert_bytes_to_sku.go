package ozon

import (
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/sku_getter"
)

func (c ozonController) ConvertBytesToSku(data *[]byte) (*[]domain.Sku, error) {
	return sku_getter.NewCSVBytesSkuGetter().GetSkus(data)
}