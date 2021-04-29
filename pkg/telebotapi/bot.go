package telebotapi

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Bot struct {
	Token string
}

func (b *Bot) getBaseUrl() (string, error) {
	if b.Token == "" {
		return "", errors.New("bot Token not provided")
	}
	return fmt.Sprintf("https://api.telegram.org/bot%v", b.Token), nil
}

func (b *Bot) makeRequestUrl(
	method string,
	params map[string]string,
) (requestUrl string, err error) {

	baseUrl, err := b.getBaseUrl()

	if err != nil {
		return "", err
	}

	requestUrl = fmt.Sprintf("%v/%v?", baseUrl, method)

	var paramStrings []string

	for paramName, paramValue := range params {
		paramString := fmt.Sprintf("%v=%v", paramName, paramValue)
		paramStrings = append(paramStrings, paramString)
	}
	requestUrl += strings.Join(paramStrings, "&")
	return
}


func (b *Bot) makeFileUrl(filePath string) (requestUrl string, err error) {
	if b.Token != "" {
		requestUrl = fmt.Sprintf("https://api.telegram.org/file/bot%v/%v", b.Token, filePath)
	} else {
		err = errors.New("bot Token not provided")
	}
	return
}

func (b *Bot) SendRequest(
	method string,
	params map[string]string,
) (resp *http.Response, err error) {
	requestUrl, err := b.makeRequestUrl(method, params)

	if err != nil {
		return
	}

	resp, err = http.Get(requestUrl)

	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		return resp, errors.New(fmt.Sprintf("%v at %v", resp.Status, method))
	}

	return
}

func (b *Bot) UpdatesGoroutine(
	updatesChan chan Update,
	errorChan chan error,
	interval time.Duration,
) {
	offset := 1

	for {
		updates, err := b.getUpdates(offset)

		for _, update := range updates {
			updatesChan <- update
			offset = update.UpdateID + 1
		}

		if err != nil {
			errorChan <- err
		}

		time.Sleep(interval)
	}
}
