package bot

import (
	"github.com/artemmarkaryan/wb-parser/internal/controller/ozon"
	t "github.com/artemmarkaryan/wb-parser/pkg/telebotapi"
	"log"
	"os"
	"time"
)

func NewBot() t.Bot {
	return t.Bot{Token: os.Getenv("BOT_TOKEN")}
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
			buff, err := ozon.NewOzonController().Process(&content)
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
