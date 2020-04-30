package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/generator"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/display"
)

const appID = "com.github.ob-algdatii-20ss.leistungsnachweis-teammaze"

const maxX = 5
const maxY = 5
const maxZ = 5

func main() {
	runtime.LockOSThread()

	log.Println("Execution path: ", os.Args[0])
	log.SetOutput(os.Stdout)

	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)

	display.FatalIfError("Could not initialize gtk.Application", err)

	_, err = application.Connect("startup", func() {
		log.Printf("Application Startup")

		rand.Seed(time.Now().UnixNano())
		furthestPoint := common.NewLocation(uint(rand.Intn(maxX)), uint(rand.Intn(maxY)), uint(rand.Intn(maxZ)))

		lab := generator.NewDepthFirstGenerator().GenerateLabyrinth(furthestPoint)

		mainWindow := display.CreateMainWindow(lab)
		mainWindow.Window.Show()
		application.AddWindow(mainWindow.Window)
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
