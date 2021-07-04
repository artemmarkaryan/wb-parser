package exporter

import (
	"context"
	"github.com/artemmarkaryan/wb-parser/internal/domain"
	"github.com/tealeg/xlsx"
	"io"
	"log"
)

type ExcelExporter struct{}

func (e ExcelExporter) Export(
	ctx context.Context,
	infoCh chan *domain.Info,
	writer io.Writer,
) (err error) {
	keys := []string{}

	f := xlsx.NewFile()
	sh, err := f.AddSheet("main")
	if err != nil {
		return err
	}

	for info := range infoCh {
		row := sh.AddRow()
		for _, key := range keys { // добавляем имеющиеся ключи
			c := row.AddCell()
			if val, hasKey := (*info)[key]; hasKey {
				c.Value = val
				info.Pop(key) // удаляем добавленные ключи из Info
			} else {
				c.Value = ""
			}
		}

		for key, value := range *info { // добавляем оставшиеся ключи
			keys = append(keys, key)
			c := row.AddCell()
			c.Value = value
		}
	}


	namesRow, addRowErr := sh.AddRowAtIndex(0)
	if addRowErr != nil {
		log.Panic(addRowErr.Error())
		return
	}
	for _, key := range keys {
		cell := namesRow.AddCell()
		cell.Value = key
	}

	log.Print("✅ Converted")

	err = f.Write(writer)
	return err
}
