package display

import (
	"errors"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"runtime"
)

// This construct courtesy of Russ Cox (https://groups.google.com/forum/#!msg/golang-nuts/IiWZ2hUuLDA/SNKYYZBelsYJ)

// OpenGL Contexts use Thread-local storage and information.
// Thus, an OpenGL Context is only valid for one specific thread.
// init is run before main.main. With the call to runtime.LockOSThread we lock the main goroutine to the main OS-Thread.
// Thus ensuring that anything called with do({...}) is run on the main thread _IF_ main.main calls display.Main()
// Also see http://github.com/golang/go/wiki/LockOSThread

//noinspection GoLinterLocal
func init() {
	runtime.LockOSThread()
}

var mainfunc = make(chan func())

// Must be called by main.main
// Does not return.
func Main() {
	for f := range mainfunc {
		f()
	}
}

func doOnMain(f func()) {
	done := make(chan bool, 1)
	mainfunc <- func() {
		f()
		done <- true
	}
	<-done
}

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
