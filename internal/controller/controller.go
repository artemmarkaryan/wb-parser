package controller

import (
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/artemmarkaryan/wb-parser/internal/interactor"
	"github.com/artemmarkaryan/wb-parser/internal/parser"
	"log"
)

func Parse() {
	skus, err := domain.GetAllSku()
	if err != nil {
		log.Fatal(err)
	}

	var infos []domain.SkuInfo
	for _, sku := range skus {
		html, err := interactor.GetHTML(sku)
		if err != nil {
			log.Fatal(err)
		}

		info, err := parser.GetInfo(html)
		if err != nil {
			log.Fatal(err)
		}

		for title, value := range info {
			infos = append(infos, domain.SkuInfo{
				SkuId: sku.GetId(),
				Title: title,
				Value: value,
			})
		}
	}

	log.Print(infos)
}