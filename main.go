package main

import (
	"github.com/artemmarkaryan/wb-parser/internal/bot"
	"github.com/artemmarkaryan/wb-parser/internal/config"
	"log"
)

func main() {
	err := config.LoadDotEnv()
	if err != nil {
		log.Panic(err.Error())
	}

	b := bot.NewBot()
	go bot.Poll(b)
	select {}
}
