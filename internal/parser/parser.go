package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"strings"
)



func GetInfo(body []byte) (infos map[string]string, err error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))

	if err != nil {
		return
	}
	infos = map[string]string{}
	doc.Find(".pp").Each(func(i int, s *goquery.Selection) {
		if s.Children().Length() == 2 {
			title :=  strings.Trim(s.Children().First().Text(), " \n")
			value := strings.Trim(s.Children().Last().Text(), " \n")
			infos[title] = value
		}
	})
	infos["Цена"] = doc.Find(".final-cost").Text()
	infos["Старая цена"] = doc.Find(".old-price").Children().Text()
	infos["Описание"] = doc.Find(".description-text").Children().Text()
	infos["Состав"] = doc.Find(".j-consist").Text()


	return
}

