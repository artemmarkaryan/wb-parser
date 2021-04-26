package excel

import (
	"github.com/tealeg/xlsx"
)

func getKeys(maps *[]map[string]string) (keys []string) {
	keyMap := make(map[string]bool)
	for _, m := range *maps {
		for k := range m {
			keyMap[k] = true
		}
	}

	for k := range keyMap {
		keys = append(keys, k)
	}
	return
}

func ConvertAndSave(maps []map[string]string, toFilePath string) (err error) {
	f := xlsx.NewFile()
	sh, err := f.AddSheet("main")
	if err != nil {
		return
	}
	colNameRow := sh.AddRow()
	keys := getKeys(&maps)
	for _, k := range keys {
		c := colNameRow.AddCell()
		c.Value = k
	}

	for _, m := range maps {
		v := sh.AddRow()
		for _, key := range keys {
			c := v.AddCell()
			if val, hasKey := m[key]; hasKey {
				c.Value = val
			} else {
				c.Value = ""
			}
		}
	}

	err = f.Save(toFilePath)
	return
}
