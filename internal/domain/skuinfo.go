package domain

type SkuInfo struct {
	SkuId string
	Title string
	Value string
}
//
//func (si *SkuInfo) Update() (err error) {
//	query := `
//UPD
//`
//}
//
//func (si *SkuInfo) Create() (err error) {
//
//}
//
//func (si *SkuInfo) UpdateOrCreate() (err error) {
//	query := `
//SELECT value FROM sku_info
//WHERE sku_id=:SkuId AND title=:Title
//`
//	db, err := NewDB()
//	if err != nil {
//		return
//	}
//	defer func() { _ = db.Close() }()
//
//	rows, err := db.NamedQuery(query, si)
//	if err != nil {
//		return
//	}
//
//	if rows.Next() {
//		var newValue string
//		err = rows.Scan(&newValue)
//		if err != nil {
//			return
//		}
//
//		err = si.Update()
//		return
//
//	} else {
//		err = si.Create()
//		return
//	}
//
//}
