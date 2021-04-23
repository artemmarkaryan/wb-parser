package mapToCSV

import (
	"encoding/csv"
	"os"
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

func ConvertMany(maps []map[string]string, filename string) (err error) {
	f, err := os.OpenFile(filename + ".csv", os.O_RDWR|os.O_CREATE, 0754)
	if err != nil {
		return
	} else {
		defer func() { _ = f.Close() }()
	}

	csvWriter := csv.NewWriter(f)
	csvWriter.Comma = '$'

	keys := getKeys(&maps)
	err = csvWriter.Write(keys)

	for _, m := range maps {
		var record []string
		for _, key := range keys {
			if val, hasKey := m[key]; hasKey {
				record = append(record, val)
			} else {
				record = append(record, "")
			}
		}
		err = csvWriter.Write(record)
	}
	csvWriter.Flush()
	return
}
