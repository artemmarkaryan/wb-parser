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

	var fromFile string
	//fmt.Println("Перенеси сюда нужный файл (.csv)")
	//_, err = fmt.Scanln(&fromFile)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	// mockup
	fromFile = `/Users/artemmarkaryan/Desktop/wb-parser-ids.csv`

	err = controller.ProcessFile(
		fromFile,
		"dumps/dump"+strconv.Itoa(int(time.Now().Unix())),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
