package sku_getter

import (
	"bytes"
	"encoding/csv"
	"errors"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"strconv"
)

// get sku instances from file content
type csvBytesSkuGetter struct {}

func NewCSVBytesSkuGetter() *csvBytesSkuGetter {
	return &csvBytesSkuGetter{}
}
func (g csvBytesSkuGetter) GetSkus(data *[]byte) (skus []domain.Sku, err error) {
	reader := csv.NewReader(bytes.NewReader(*data))
	reader.Comma = ';'
	skuStrings, err := reader.ReadAll()
	if err != nil {
		return
	}

	for _, record := range skuStrings {
		if len(record) != 2 {
			return nil, errors.New("В строке должно быть ровно значения, а сейчас " + strconv.Itoa(len(record)))
		}
		skus = append(skus, domain.Sku{
			Id:  record[0],
			Url: record[1],
		})
	}

	return
}
