package bot

import (
	"github.com/artemmarkaryan/wb-parser/internal/config"
	"github.com/artemmarkaryan/wb-parser/internal/controller/ozon"
	"log"
	"testing"
)

func TestOzon(t *testing.T) {
	if err := config.LoadDotEnv(); err != nil {
		t.Error(err.Error())
	}

	b := NewBot()
	content, err := b.GetFileContent("documents/file_63.csv")
	if err != nil {
		log.Panic(err)
	}

	buff, err := ozon.NewOzonController().Process(&content)
	if err != nil {
		log.Panic(err)
	}

	err = b.SendFile(296286381, buff)
	if err != nil {
		log.Panic(err)
	}
}
