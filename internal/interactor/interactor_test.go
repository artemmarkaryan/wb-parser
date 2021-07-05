package interactor

import (
	"github.com/artemmarkaryan/wb-parser/internal/config"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"testing"
)


func TestGetHTML(t *testing.T) {
	_ = config.LoadDotEnv()
	skus, err := domain.GetAllSku()
	if err != nil {t.Error(err.Error())}

	for _, sku := range skus {
		body, err := GetHTML(sku)
		if err != nil {t.Error(err.Error())}
		t.Log(body.Read([]byte{}))
	}
}
