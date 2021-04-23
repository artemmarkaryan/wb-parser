package main

import (
	"github.com/artemmarkaryan/wb-parser/internal/config"
	"github.com/artemmarkaryan/wb-parser/internal/controller"
	"log"
	"strconv"
	"time"
)

func main() {
	err := config.LoadDotEnv()
	if err != nil {
		log.Fatal(err.Error())
	}

	controller.Parse("dumps/dump"+strconv.Itoa(int(time.Now().Unix()))+".json")

}
