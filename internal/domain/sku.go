package domain

import "strings"

type Sku struct {
	Id  string
	Url string
}

func (s *Sku) GetId() string { return s.Id }
func (s *Sku) GetUrl() string { return strings.Trim(s.Url, " \n\r") }


func GetAllSku() (skus []Sku, err error) {
	db, err := NewDB()
	if err != nil {
		return
	}
	defer func() { _ = db.Close() }()

	query := "select * from sku"
	rows, err := db.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		newSku := Sku{}
		err = rows.Scan(&newSku.Id, &newSku.Url)
		if err != nil {
			return
		}
		skus = append(skus, newSku)
	}
	return
}
