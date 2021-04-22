package controller

import (
	"encoding/json"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/interactor"
	"github.com/artemmarkaryan/wb-parser/internal/parser"
	"log"
	"os"
)

func Parse(toFile string) {

	err := os.WriteFile(
		toFile,
		[]byte{},
		0666,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	skus, err := domain.GetAllSku()
	if err != nil {
		log.Fatal(err)
	}

	var infos []map[string]string

	for _, sku := range skus {
		html, err := interactor.GetHTML(sku)
		if err != nil {
			log.Fatal(err.Error())
		}

		info, err := parser.GetInfo(html)
		if err != nil {
			log.Fatal(err.Error())
		}

		info["id"] = sku.GetId()
		info["url"] = sku.GetUrl()
		infos = append(infos, info)
	}

	infoJSON, err := json.Marshal(infos)

	err = os.WriteFile(
		toFile,
		infoJSON,
		0666,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

}
