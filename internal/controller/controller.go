package controller

import (
	"bytes"
	"errors"
	d "github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/interactor"
	"github.com/artemmarkaryan/wb-parser/internal/parser"
	"github.com/artemmarkaryan/wb-parser/pkg/excel"
	"github.com/artemmarkaryan/wb-parser/pkg/make-http-client"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const poolSize = 10
const coolDown = time.Second / 50

type Info map[string]string

func (i *Info) Map() map[string]string {
	return *i
}

func ProcessFile(fromFile, toFile string) (err error) {
	if err = checkFile(fromFile); err != nil {
		return errors.New("Невозможно открыть файл: " + err.Error())
	}

	getter, err := defineFileGetter(fromFile)
	if err != nil {
		return
	}

	writeToFile(parse(getter), toFile)
	return
}

func ProcessData(data *[]byte) (buff *bytes.Buffer, err error) {
	getter := interactor.NewCSVBytesSkuGetter(*data)
	infoArr := parse(getter)
	var processedData []byte
	buff = bytes.NewBuffer(processedData)
	err = excel.ConvertAndWrite(infoArrToMap(infoArr), buff)
	return
}

// check if from exists
func checkFile(file string) error {
	f, err := os.OpenFile(file, os.O_RDONLY, 0777)
	if err != nil {
		return err
	} else {
		_ = f.Close()
		return nil
	}
}

// match getter by extension
func defineFileGetter(file string) (getter interactor.SkuGetter, err error) {
	switch filepath.Ext(file) {
	case ".csv":
		getter = interactor.NewCSVFileSkuGetter(file)
	default:
		return nil, errors.New("unknown file extension")
	}
	return
}

func parse(getter interactor.SkuGetter) (infoArr []Info) {
	//var pool []func(chan d.Sku, chan Info, chan error)
	allSku, _ := getter.GetSkus()
	skuCh := make(chan d.Sku, len(allSku))
	infoCh := make(chan Info, len(allSku))
	errCh := make(chan error, len(allSku))
	client := makeHTTPClient.NewHTTPClient(len(allSku))
	go func() {
		for _, s := range allSku {
			skuCh <- s
		}
	}() // put sku in channel

	for i := 0; i < poolSize; i++ {
		num := i
		go func(sCh chan d.Sku, iCh chan Info, eCh chan error) {
			for sku := range sCh {
				makeRequest(sku, client, iCh, eCh)
				log.Printf("goroutine #%v: sku #%v", num, sku.GetId())
				time.Sleep(coolDown)
			}
		} (skuCh, infoCh, errCh)
	} // get sku from channel; put to result channels

	var rcv int
	// read from result channel
	for rcv < len(allSku) {
			select {
			case i := <-infoCh:
				infoArr = append(infoArr, i)
				rcv++
			case e := <-errCh:
				log.Print(e.Error())
				rcv++
			}
		}

	return
}

func writeToFile(infoArr []Info, path string) {
	_ = excel.ConvertAndSave(
		infoArrToMap(infoArr),
		path,
	)
}

func infoArrToMap(infoArr []Info) (ms []map[string]string) {
	for _, info := range infoArr {
		ms = append(ms, info)
	}
	return
}

func makeRequest(sku d.Sku, client *http.Client, iCh chan Info, eCh chan error) {
	body, err := interactor.GetHTML(sku, client)
	if err != nil {
		eCh <- err
		return
	}

	info, err := parser.GetInfo(body)
	info["id"] = sku.GetId()
	info["url"] = sku.GetUrl()

	if err != nil {
		eCh <- err
		return
	} else {
		info["id"] = sku.GetId()
		info["url"] = sku.GetUrl()
		iCh <- info
	}

	return
}

//func splitSkuArr(input []d.Sku, size int) (result [][]d.Sku) {
//	n := len(input)
//	if len(input) <= size {
//		result = append(result, input)
//		return
//	}
//	result = append(result, splitSkuArr(input[:n/2], size)...)
//	result = append(result, splitSkuArr(input[n/2:], size)...)
//	return
//}
