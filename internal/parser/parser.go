package parser

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"io"
	"strings"
)



func GetInfo(body io.Reader) (infos map[string]string, err error) {
	_, err = html.Parse(body)
	if err != nil {
		return nil, errors.New("cant parse body: " + err.Error())
	}

	doc, err := goquery.NewDocumentFromReader(body)
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

