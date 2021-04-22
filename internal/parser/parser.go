package parser

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"strings"
)



func GetInfo(body io.ReadCloser) (infos map[string]string, err error) {
	doc, err := goquery.NewDocumentFromReader(body)
	defer func() {_ = body.Close()}()
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
	return
}

