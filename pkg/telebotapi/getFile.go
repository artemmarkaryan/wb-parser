package telebotapi

import (
	"encoding/json"
	"io"
	"log"
)

func (b *Bot) GetFile(fileId string) (f FileResponse, err error) {
	resp, err := b.SendRequest(
		"getFile",
		map[string]string{"file_id": fileId},
	)
	if err != nil {
		return
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	f = FileResponse{}
	
	log.Print(string(body))
	err = json.Unmarshal(body, &f)

	return
}
