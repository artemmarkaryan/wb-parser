package main

import (
	"github.com/artemmarkaryan/wb-parser/internal/bot"
	"github.com/artemmarkaryan/wb-parser/internal/config"
	"log"
)

func main() {
	log.Print("Running")
	defer log.Print("Stopping")

	err := config.LoadDotEnv()
	if err != nil {
		log.Println(err.Error())
	}

	b := bot.NewBot()
	go bot.Poll(b)
	select {}
}
