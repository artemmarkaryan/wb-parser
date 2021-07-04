package wildberries

import (
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/sku_getter"
)

func (w wildberriesController) ConvertBytesToSku(data *[]byte) (*[]domain.Sku, error) {
	return sku_getter.NewCSVBytesSkuGetter().GetSkus(data)
}