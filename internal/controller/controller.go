package controller

import (
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
	"strings"
	"sync"
	"time"
)

const buffSizeUpperBound = 10
const coolDown = time.Second / 10

type Info map[string]string

func (i *Info) Map() map[string]string {
	return *i
}

func ProcessFile(fromFile, toFile string) (err error) {
	if err = checkFile(fromFile); err != nil {
		return errors.New("Невозможно открыть файл: " + err.Error())
	}

	getter, err := defineGetter(fromFile)
	if err != nil {
		return
	}

	err = parse(toFile, getter)
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
func defineGetter(file string) (getter interactor.SkuGetter, err error) {
	switch filepath.Ext(file) {
	case ".csv":
		getter = interactor.NewCSVSkuGetter(file)
	default:
		return nil, errors.New("unknown file extension")
	}
	return
}

func parse(toFile string, getter interactor.SkuGetter) (err error) {
	var infos []Info
	var fails []d.Sku
	allSku, err := getter.GetSkus()
	fails, err = processSkuArr(allSku, &infos)
	unfixed, err := processSkuArr(fails, &infos)
	if err != nil {
		return
	}
	if len(unfixed) > 0 {
		return errors.New("Остались необработанными: " + func() string {
			var urls []string
			for _, sku := range unfixed {
				urls = append(urls, sku.GetUrl())
			}
			return strings.Join(urls, "\n")
		}())
	}

	func() (ms []map[string]string) {
		for _, info := range infos {
			ms = append(ms, info)
		}
		return
	}()

	err = excel.ConvertAndSave(
		// convert infos to map[string]string
		func() (ms []map[string]string) {
			for _, info := range infos {
				ms = append(ms, info)
			}
			return
		}(),
		toFile,
	)
	if err != nil {
		return
	}

	return
}

func processSkuArr(skuArr []d.Sku, infos *[]Info) (fails []d.Sku, err error) {
	allSkuBuffered := splitSkuArr(skuArr, buffSizeUpperBound)

	wg := sync.WaitGroup{}

	for i, buffer := range allSkuBuffered {
		buffSize := len(buffer)
		time.Sleep(coolDown)
		log.Printf("Processing values %v - %v", i*buffSize+1, (i+1)*buffSize)

		infoChan := make(chan Info, buffSize)
		failChan := make(chan d.Sku, buffSize)
		errChan := make(chan error, buffSize)

		httpClient := makeHTTPClient.NewHTTPClient(buffSize)
		for _, sku := range buffer {
			wg.Add(1)
			go makeRequest(sku, &wg, httpClient, infoChan, failChan, errChan)
		}

		wg.Wait()
		close(infoChan)
		close(failChan)
		close(errChan)

		for info := range infoChan {
			*infos = append(*infos, info)
			log.Printf("received from %v", info["url"])
		}
		for fail := range failChan {
			fails = append(fails, fail)
			log.Printf("%v failed", fail.GetUrl())
		}
		for err := range errChan {
			log.Print(err.Error())
		}
	}

	return
}

func makeRequest(
	sku d.Sku,
	wg *sync.WaitGroup,
	httpClient *http.Client,
	infoChan chan Info,
	failChan chan d.Sku,
	errChan chan error,
) {
	defer wg.Done()
	body, err := interactor.GetHTML(sku, httpClient)
	if err != nil {
		errChan <- err
	}

	info, err := parser.GetInfo(body)
	info["id"] = sku.GetId()
	info["url"] = sku.GetUrl()

	if err != nil {
		failChan <- sku
		errChan <- err
	} else {
		info["id"] = sku.GetId()
		info["url"] = sku.GetUrl()
		infoChan <- info
	}
}

func splitSkuArr(input []d.Sku, size int) (result [][]d.Sku) {
	n := len(input)
	if len(input) <= size {
		result = append(result, input)
		return
	}
	result = append(result, splitSkuArr(input[:n/2], size)...)
	result = append(result, splitSkuArr(input[n/2:], size)...)
	return
}
