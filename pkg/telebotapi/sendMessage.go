package telebotapi

import "strconv"

func (b *Bot) SendMessage(chatId int, text string) (err error){
	_, err = b.SendRequest(
		"sendMessage",
		map[string]string{
			"chat_id": strconv.Itoa(chatId),
			"text": text,
		})
	return
}
