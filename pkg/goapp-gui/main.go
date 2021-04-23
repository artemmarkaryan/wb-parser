package goappGUI

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	_ "github.com/maxence-charriere/go-app/v8/pkg/app"
	"log"
	"net/http"
)

type home struct {
	app.Compo
}

func (h *home) Render() app.UI {
	return app.H1().Text("Hello World!")
}

func Demo() {
	app.Route("/", &home{})
	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
