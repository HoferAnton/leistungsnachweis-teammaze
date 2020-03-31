package display

import (
	"errors"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func FatalIfError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func asWindow(obj glib.IObject) (*gtk.ApplicationWindow, error) {
	if win, ok := obj.(*gtk.ApplicationWindow); ok {
		return win, nil
	}

	return nil, errors.New("not a *gtk.ApplicationWindow")
}
