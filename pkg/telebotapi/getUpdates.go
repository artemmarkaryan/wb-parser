package telebotapi

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func (b *Bot) getUpdates(offset int) (updates []Update, err error) {
	requestUrl, err := b.makeRequestUrl(
		"getUpdates",
		map[string]string{"offset": strconv.Itoa(offset)},
	)

	if err != nil {
		return
	}

	resp, err := http.Get(requestUrl)

	if err != nil {
		return
	}

	updates, err = ParseUpdateResponse(*resp)
	return
}

func ParseUpdateResponse(httpResponse http.Response) (updates []Update, err error) {
	body, err := io.ReadAll(httpResponse.Body)

	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	updates = response.Result
	return
}

