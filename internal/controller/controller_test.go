package controller

import (
	"github.com/artemmarkaryan/wb-parser/internal/config"
	"testing"
)

func TestParse(t *testing.T) {
	_ = config.LoadDotEnv()
	Parse()
}
