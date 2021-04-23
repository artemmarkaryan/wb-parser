package main

import (
	"fmt"
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

	var fromFile string
	_, err = fmt.Scanln(fromFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = controller.ProcessFile(
		fromFile,
		"dumps/dump"+strconv.Itoa(int(time.Now().Unix())),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
