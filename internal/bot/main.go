package bot

import (
	"github.com/artemmarkaryan/wb-parser/internal/controller"
	t "github.com/artemmarkaryan/wb-parser/pkg/telebotapi"
	"log"
	"time"
)

func NewBot() t.Bot {
	return t.Bot{Token: "1780566572:AAGo452wn9oot8AwMNW4q8KVYF9SdfXndJ0"}
}

func Poll(b t.Bot) {
	uCh := make(chan t.Update)
	eCh := make(chan error)

	go b.UpdatesGoroutine(uCh, eCh, time.Second/20)

	for {
		select {
		case upd := <-uCh:
			if upd.Message.Document.FileId == "" {
				err := b.SendMessage(upd.Message.Chat.ID, "😎")
				if err != nil {
					log.Print(err)
				}
			}
			f, err := b.GetFile(upd.Message.Document.FileId)
			if err != nil {
				log.Print(err.Error())
				break
			}

			content, err := b.GetFileContent(f.Result.FilePath)
			if err != nil {
				log.Print(err.Error())
				break
			}

			_ = b.SendMessage(upd.Message.Chat.ID, "Начал обработку")
			buff, err := controller.ProcessData(&content)
			if err != nil {
				_ = b.SendMessage(
					upd.Message.Chat.ID,
					"Ошибка во время сбора данных: " + err.Error(),
				)
				log.Print(err.Error())
				break
			}

			err = b.SendFile(upd.Message.Chat.ID, buff)
			if err != nil {
				log.Print(err.Error())
				break
			}

		case err := <-eCh:
			log.Print(err.Error())
		}
	}
}
