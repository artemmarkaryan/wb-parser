package main

import (
	"fmt"
	"github.com/artemmarkaryan/wb-parser/internal/config"
	"github.com/artemmarkaryan/wb-parser/internal/controller"
	"log"
	"path/filepath"
	"time"
)

func main() {
	err := config.LoadDotEnv()
	if err != nil {
		log.Fatal(err.Error())
	}

	var fromFile string
	fmt.Println("Перенеси сюда нужный файл (.csv)")
	_, err = fmt.Scanln(&fromFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	// mockup
	//fromFile = `/Users/artemmarkaryan/Desktop/wb-parser-ids.csv`
	toFileDir := filepath.Dir(fromFile)
	toFile := filepath.Join(
		toFileDir,
		time.Now().Format("02-01_15:04:05") + ".xlsx",
		//+ filepath.Ext(fromFile),
		)

	err = controller.ProcessFile(fromFile, toFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	time.Sleep(10 * time.Minute)
}
