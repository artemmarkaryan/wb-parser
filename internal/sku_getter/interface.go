package sku_getter

import "github.com/artemmarkaryan/wb-parser/internal/domain"

// SkuGetter gets sku instances from some source
type SkuGetter interface {
	GetSkus(*[]byte) ([]domain.Sku, error)
}
