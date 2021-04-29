package telebotapi

import (
	"io"
	"net/http"
)

func (b *Bot) GetFileContent(tgFilePath string) (content []byte, err error) {
	url, err := b.makeFileUrl(tgFilePath)
	if err != nil {
		return
	}

	resp, err := http.Get(url)
	defer func() { _ = resp.Body.Close() }()
	if err != nil {
		return
	}

	return io.ReadAll(resp.Body)
}
