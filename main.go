package main

import (
	"github.com/artemmarkaryan/wb-parser/internal/config"
	"log"
)

func main() {
	err := config.LoadDotEnv()
	if err != nil {
		log.Fatal(err.Error())
	}

}
