package config

import (
	"errors"
	dotenv "github.com/joho/godotenv"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func LoadDotEnv() (err error) {
	_ = dotenv.Load(basepath + "/.env")
	if os.Getenv("DATABASE_PASSWORD") == ""{
		err = errors.New("env not loaded")
	}
	return
}

