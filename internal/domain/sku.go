package domain

import "strings"

type Sku struct {
	Id  string
	Url string
}

func (s *Sku) GetId() string { return s.Id }
func (s *Sku) GetUrl() string { return strings.Trim(s.Url, " \n\r") }
