package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"strings"
)



func GetInfo(body []byte) (infos map[string]string, err error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	//_ = os.WriteFile(fmt.Sprintf("./logs/temp%v.html", time.Now().Unix()), body, 0777)
	if err != nil {
		return
	}
	infos = map[string]string{}
	doc.Find(".product-params__row").Each(func(i int, s *goquery.Selection) {
		if s.Children().Length() == 2 {
			title :=  strings.Trim(s.Children().First().Text(), " \n")
			value := strings.Trim(s.Children().Last().Text(), " \n")
			infos[title] = value
		}
	})
	infos["Название"] = doc.Find("[data-link=\"text{:productCard^goodsName}\"]").Text()
	infos["Цена"] = doc.Find(".price-block__final-price").Text()
	infos["Старая цена"] = doc.Find(".price-block__old-price").Children().Text()
	infos["Описание"] = doc.Find(".j-description").Children().Text()
	infos["Состав"] = doc.Find(".j-consist").Text()


	return
}

