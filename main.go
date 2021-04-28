package main

import (
	"fmt"
	"github.com/artemmarkaryan/wb-parser/internal/controller"
	"log"
	"path/filepath"
	"time"
)

func main() {

	var err error
	var fromFile string

	fmt.Println("Перенеси сюда нужный файл (.csv)")
	_, err = fmt.Scanln(&fromFile)
	if err != nil {
		log.Panic(err.Error())
	}
	toFileDir := filepath.Dir(fromFile)
	toFile := filepath.Join(
		toFileDir,
		time.Now().Format("15-04") + ".xlsx",
		//+ filepath.Ext(fromFile),
		)

	timeF := time.Now()

	err = controller.ProcessFile(fromFile, toFile)
	if err != nil {
		log.Panic(err.Error())
	}

	log.Printf("Процесс занял %v", time.Now().Sub(timeF))
}
