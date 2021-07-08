package main

import (
	"github.com/artemmarkaryan/wb-parser/internal/bot"
	"github.com/artemmarkaryan/wb-parser/internal/config"
	"log"
)

func main() {
	log.Print("Running")
	err := config.LoadDotEnv()
	if err != nil {
		log.Println(err.Error())
	}

	b := bot.NewBot()
	go bot.Poll(b)
	select {}

	//log.Print(controller.ProcessFile(
	//	"/Users/artemmarkaryan/Downloads/l.csv",
	//	"/Users/artemmarkaryan/Downloads/q3123423.csv",
	//))
}
