package controller

import (
	d "github.com/artemmarkaryan/wb-parser/internal/domain"
	"testing"
)

func TestSplitSkuArr(t *testing.T) {
	skuArr := []d.Sku{
		{Id:  "1"},
		{Id:  "2"},
		//{Id:  "3"},
		{Id:  "4"},
	}
	r := splitSkuArr(skuArr, 2)
	t.Log(r)
}
