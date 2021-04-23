package controller

import (
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/interactor"
	"github.com/artemmarkaryan/wb-parser/internal/parser"
	"github.com/artemmarkaryan/wb-parser/pkg/make-http-client"
	"github.com/artemmarkaryan/wb-parser/pkg/map-to-csv"
	"log"
	"time"
)

func Parse(filename string) {
	skus, err := domain.GetAllSku()
	if err != nil {
		log.Fatal(err)
	}

	infoChan := make(chan map[string]string, len(skus))

	for _, sku := range skus {
		httpClient := makeHTTPClient.NewHTTPClient(len(skus))
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

			err = mapToCSV.ConvertMany(infos, filename)
			//infoJSON, err := json.Marshal(infos)
			//
			//err = os.WriteFile(
			//	filename,
			//	infoJSON,
			//	0666,
			//)

			if err != nil {
				log.Fatal(err.Error())
			}

			break
		}
	}
}

