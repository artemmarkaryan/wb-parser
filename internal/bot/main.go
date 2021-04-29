package bot

import (
	"github.com/artemmarkaryan/wb-parser/internal/controller"
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
				continue
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

			buff, err := controller.ProcessData(&content)
			if err != nil {
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
