package fyneGUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
)

func Demo() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.Size{
		Width:  500,
		Height: 400,
	})

	buttonWidget := widget.NewButton(
		"Выбрать файл",
		func() {
			dialog.ShowFileOpen(
				func(closer fyne.URIReadCloser, err error) {
					var content []byte
					_, _ = closer.Read(content)
					log.Print(string(content))
					
					_ = closer.Close()
				},
				w,
			)
		},
	)

	w.SetContent(buttonWidget)

	w.ShowAndRun()
}
