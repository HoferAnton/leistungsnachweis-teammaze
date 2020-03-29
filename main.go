package main

import (
	"errors"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/display"
	"log"
	"os"
)

const appId = "com.github.ob-algdatii-20ss.leistungsnachweis-teammaze"

func main() {

	log.Println(os.Args[0])

	log.SetOutput(os.Stdout)
	application, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)

	display.FatalIfError("Could not initialize gtk.Application", err)

	_, err = application.Connect("startup", func() {
		log.Printf("Application Startup")
	})

	display.FatalIfError("Startup Signal Connection failed: ", err)

	_, err = application.Connect("activate", func() {
		log.Print("Application Activate")

		builder, err := gtk.BuilderNewFromFile("display/ui/glarea.ui")

		display.FatalIfError("Could not create GTK Builder: ", err)

		signals := map[string]interface{}{
			"gl_init": display.Realize,   // Called on Window Creation
			"gl_draw": display.Render,    // Window Redraw
			"gl_fini": display.Unrealize, // Window Deletion
		}

		builder.ConnectSignals(signals)

		obj, err := builder.GetObject("main_window")

		display.FatalIfError("Could not get main_window: ", err)

		win, err := isWindow(obj)

		display.FatalIfError("Could not use main_window: ", err)

		win.Show()

		application.AddWindow(&(win.Window))
	})

	display.FatalIfError("Activation Signal Connection failed: ", err)

	_, err = application.Connect("shutdown", func() {
		log.Println("Application Shutdown!")
	})

	display.FatalIfError("Shutdown Signal Connection failed: ", err)

	os.Exit(application.Run(os.Args))
}

func isWindow(obj glib.IObject) (*gtk.ApplicationWindow, error) {
	if win, ok := obj.(*gtk.ApplicationWindow); ok {
		return win, nil
	}

	return nil, errors.New("not a *gtk.ApplicationWindow")
}
