package interactor

import "github.com/artemmarkaryan/wb-parser/internal/domain"

type Sku struct {
	Id  int
	Url string
}

type SkuInfo struct {
	SkuId int
	Title string
	Value string
}

func GetAllSku() (skus []Sku, err error) {
	db, err := domain.NewDB()
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
