package interactor

import (
	"github.com/artemmarkaryan/wb-parser/internal/config"
	"testing"
)

func TestGetAllSku(t *testing.T) {
	if err := config.LoadDotEnv(); err != nil {
		t.Error(err.Error())
	}
	t.Log(GetAllSku())
}