package telebotapi

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
)

func (b *Bot) SendFile(chatId int, data io.Reader) (err error) {
	url, err := b.makeRequestUrl(
		"sendDocument",
		map[string]string{
			"chat_id": strconv.Itoa(chatId),
		},
	)
	buff := new(bytes.Buffer)
	w := multipart.NewWriter(buff)
	fileWriter, err := w.CreateFormFile("document", "result.xlsx")
	if err != nil {
		return
	}

	written, err := io.Copy(fileWriter, data)
	if err != nil {
		return
	} else {
		log.Printf("written %v bytes", written)
	}

	err = w.Close()
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, url, buff)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	_, err = io.ReadAll(resp.Body)
	return
}

