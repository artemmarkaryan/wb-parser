package controller

import (
	"errors"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/interactor"
	"github.com/artemmarkaryan/wb-parser/internal/parser"
	"github.com/artemmarkaryan/wb-parser/pkg/make-http-client"
	"github.com/artemmarkaryan/wb-parser/pkg/map-to-csv"
	"log"
	"os"
	"path/filepath"
	"time"
)

func ProcessFile(fromFile, toFile string) (err error) {
	// check if fromFile exists
	f, err := os.OpenFile(fromFile, os.O_RDONLY, 0777)
	if err != nil {
		return
	} else {
		_ = f.Close()
	}

	// match getter by extension
	extension := filepath.Ext(fromFile)
	var getter interactor.SkuGetter
	switch extension {
	case ".csv":
		getter = interactor.NewCSVSkuGetter(fromFile)
	default:
		return errors.New("unknown file extension")
	}

	// call parser
	err = parse(toFile, getter)
	if err != nil {
		return
	}
	return
}

func parse(toFile string, getter interactor.SkuGetter) (err error) {
	skus, err := getter.GetSkus()
	if err != nil {
		return
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
				return
			}

			info, err := parser.GetInfo(html)
			if err != nil {
				return
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

			err = mapToCSV.ConvertMany(infos, toFile)
			if err != nil {
				return
			}

			break
		}
	}
	return
}
