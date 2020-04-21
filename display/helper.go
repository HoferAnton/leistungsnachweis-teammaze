package display

import (
	"errors"
	"log"

	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func FatalIfError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func checkForGLError(message string) {
	err := gl.GetError()
	if err != gl.NO_ERROR {
		log.Fatalf("%s ErrorCode: 0x%X", message, err)
	}
}

func asWindow(obj glib.IObject) (*gtk.ApplicationWindow, error) {
	if win, ok := obj.(*gtk.ApplicationWindow); ok {
		return win, nil
	}

	return nil, errors.New("not a *gtk.ApplicationWindow")
}
