package excel

import (
	"errors"
	"github.com/tealeg/xlsx"
	"io"
	"log"
)

type Map interface {
	Range() []interface{}
	Get(interface{}) interface{}
}

func getKeys(maps []map[string]string) (keys []string) {
	keyMap := make(map[string]bool)
	for _, m := range maps {
		for k := range m {
			keyMap[k] = true
		}
	}

	for k := range keyMap {
		keys = append(keys, k)
	}
	return
}

func Convert(maps []map[string]string) (f *xlsx.File, err error) {
	f = xlsx.NewFile()
	sh, err := f.AddSheet("main")
	if err != nil {
		return
	}
	colNameRow := sh.AddRow()
	keys := getKeys(maps)
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
	log.Print("✅ Converted")
	return
}

func ConvertAndWrite(maps []map[string]string, writer io.Writer) (err error) {
	f, err := Convert(maps)
	if err != nil {
		return
	}

	err = f.Write(writer)
	if err != nil {
		return
	}

	log.Print("✅ Wrote to buffer")
	return
}

func ConvertAndSave(maps []map[string]string, toFilePath string) (err error) {
	f, err := Convert(maps)
	if err != nil {
		return
	}

	err = f.Save(toFilePath)
	if err != nil {
		err = errors.New("cant save .xlsx file: " + err.Error())
	}
	return
}
