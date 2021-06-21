package controller

import (
	"bytes"
	d "github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/html_retriever"
	"github.com/artemmarkaryan/wb-parser/internal/parser"
	"github.com/artemmarkaryan/wb-parser/internal/sku_getter"
	"github.com/artemmarkaryan/wb-parser/pkg/excel"
	makeHTTPClient "github.com/artemmarkaryan/wb-parser/pkg/make-http-client"
	"log"
	"net/http"
	"time"
)

type ControllerParent struct {
	poolSize  uint
	coolDown  time.Duration
	skuGetter *sku_getter.SkuGetter
}

func NewControllerParent(
	getter sku_getter.SkuGetter,
	poolSize uint,
	coolDown time.Duration,
) *ControllerParent {
	return &ControllerParent{skuGetter: &getter, poolSize: poolSize, coolDown: coolDown}
}

func (r *ControllerParent) ProcessBytes(data *[]byte) (*bytes.Buffer, error) {
	infoArr := r.parse(data)
	var processedData []byte

	buff := bytes.NewBuffer(processedData)

	var infoMaps []map[string]string
	for _, info := range infoArr {
		infoMaps = append(infoMaps, info)
	}

	err := excel.ConvertAndWrite(infoMaps, buff)
	return buff, err
}

func (r *ControllerParent) parse(data *[]byte) (infoArr []Info) {
	allSku, _ := (*r.skuGetter).GetSkus(data)
	skuCh := make(chan d.Sku, len(allSku))
	infoCh := make(chan Info, len(allSku))
	errCh := make(chan error, len(allSku))
	client := makeHTTPClient.NewHTTPClient(len(allSku))
	go func() {
		for _, s := range allSku {
			skuCh <- s
		}
	}() // put sku in channel

	for i := uint(0); i < r.poolSize; i++ {
		num := i
		go func(sCh chan d.Sku, iCh chan Info, eCh chan error) {
			for sku := range sCh {
				makeRequest(sku, client, iCh, eCh)
				log.Printf("goriutine #%v: sku #%v", num, sku.GetId())
				time.Sleep(r.coolDown)
			}
		}(skuCh, infoCh, errCh)
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

func makeRequest(sku d.Sku, client *http.Client, iCh chan Info, eCh chan error) {
	body, err := html_retriever.WildberriesHtmlRetriever{}.GetHTML(sku, client)
	if err != nil {
		eCh <- err
		return
	}

	info, err := parser.WildberriesParser{}.GetInfo(body)
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
