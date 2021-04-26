package controller

import (
	"errors"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/interactor"
	"github.com/artemmarkaryan/wb-parser/internal/parser"
	"github.com/artemmarkaryan/wb-parser/pkg/excel"
	"github.com/artemmarkaryan/wb-parser/pkg/make-http-client"
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
	log.Print("finished parsing")
	return
}

func parse(toFile string, getter interactor.SkuGetter) (err error) {
	skus, err := getter.GetSkus()
	if err != nil {
		return
	}

	var infos []map[string]string
	infoChan := make(chan map[string]string, len(skus))
	errChan := make(chan error, len(skus))

	for _, sku := range skus {
		httpClient := makeHTTPClient.NewHTTPClient(len(skus))
		go func(sku domain.Sku) {
			html, err := interactor.GetHTML(
				sku,
				httpClient,
			)
			if err != nil {
				errChan <- err
			}

			info, err := parser.GetInfo(html)
			if err != nil {
				errChan <- err
			}

			info["id"] = sku.GetId()
			info["url"] = sku.GetUrl()
			infoChan <- info
		}(sku)
	}

	for {
		time.Sleep(time.Second/50)
		select {
		case err := <- errChan:
			log.Print(err.Error())
		case info := <- infoChan:
			log.Printf("recieved from %v", info["url"])
			infos = append(infos, info)
		}
		if len(infos) == len(skus) {
			err = excel.ConvertAndSave(infos, toFile)
			if err != nil {
				return
			}

			break
		}
	}
	return
}
