package main

import (
	"log"
	"os"
	"runtime"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/display"
)

const appID = "com.github.ob-algdatii-20ss.leistungsnachweis-teammaze"

func main() {
	runtime.LockOSThread()

	log.Println(os.Args[0])

	log.SetOutput(os.Stdout)
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)

	display.FatalIfError("Could not initialize gtk.Application", err)

	_, err = application.Connect("startup", func() {
		log.Printf("Application Startup")

		mainWindow := display.CreateMainWindow()

		mainWindow.Show()
		application.AddWindow(mainWindow)
	})
	display.FatalIfError("Startup Signal Connection failed: ", err)
	_, err = application.Connect("activate", func() {
		log.Print("Application Activate")
	})
	display.FatalIfError("Activation Signal Connection failed: ", err)
	_, err = application.Connect("shutdown", func() {
		log.Println("Application Shutdown!")
	})
	display.FatalIfError("Shutdown Signal Connection failed: ", err)
	os.Exit(application.Run(os.Args))
}
