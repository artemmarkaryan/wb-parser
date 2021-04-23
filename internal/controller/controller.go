package controller

import (
	"encoding/json"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/interactor"
	"github.com/artemmarkaryan/wb-parser/internal/parser"
	make_http_client "github.com/artemmarkaryan/wb-parser/pkg/make-http-client"
	"log"
	"os"
	"time"
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

	infoChan := make(chan map[string]string, len(skus))

	for _, sku := range skus {
		httpClient := make_http_client.MakeHttpClient(len(skus))
		go func(sku domain.Sku) {
			html, err := interactor.GetHTML(
				sku,
				httpClient,
			)
			if err != nil {
				log.Fatal(err.Error())
			}

			info, err := parser.GetInfo(html)
			if err != nil {
				log.Fatal(err.Error())
			}

			info["id"] = sku.GetId()
			info["url"] = sku.GetUrl()
			infoChan <- info
		}(sku)
	}
	for {
		time.Sleep(time.Second)
		log.Printf("Recieved %v of %v", len(infoChan), cap(infoChan))
		if len(infoChan) == cap(infoChan) {
			close(infoChan)
			var infos []map[string]string
			for info := range infoChan {
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

			break
		}
	}
}
