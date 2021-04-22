package domain

import (
	"github.com/artemmarkaryan/wb-parser/internal/config"
	"testing"
)

func TestNewDB(t *testing.T) {
	if err := config.LoadDotEnv(); err != nil {
		t.Error(err)
	}

	db, err := NewDB()
	if err != nil {
		t.Error(err.Error())
		return
	}

	if err := db.Ping(); err != nil {
		t.Error(err.Error())
	}
	defer func () {_ = db.Close()}()
}