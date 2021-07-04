package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"strings"
)

type WildberriesParser struct {}

func (r WildberriesParser) GetInfo(body *domain.HtmlBody) (info domain.Info, err error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(*body))

	if err != nil {
		return
	}
	doc.Find(".pp").Each(func(i int, s *goquery.Selection) {
		if s.Children().Length() == 2 {
			title :=  strings.Trim(s.Children().First().Text(), " \n")
			value := strings.Trim(s.Children().Last().Text(), " \n")
			info[title] = value
		}
	})
	info["Цена"] = doc.Find(".final-cost").Text()
	info["Старая цена"] = doc.Find(".old-price").Children().Text()
	info["Описание"] = doc.Find(".description-text").Children().Text()
	info["Состав"] = doc.Find(".j-consist").Text()

	return
}

