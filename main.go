package main

import (
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/display"
)

const appID = "com.github.ob-algdatii-20ss.leistungsnachweis-teammaze"

const labyrinthSize = 5
const numConnections = 100

func main() {
	runtime.LockOSThread()

	log.Println(os.Args[0])

	log.SetOutput(os.Stdout)

	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)

	display.FatalIfError("Could not initialize gtk.Application", err)

	_, err = application.Connect("startup", func() {
		log.Printf("Application Startup")

		lab := common.NewLabyrinth(common.NewLocation(labyrinthSize, labyrinthSize, labyrinthSize))

		rand.Seed(time.Now().UnixNano())

		for i := 0; i < numConnections; i++ {
			randX := uint(rand.Intn(labyrinthSize))
			randY := uint(rand.Intn(labyrinthSize))
			randZ := uint(rand.Intn(labyrinthSize))

			randLoc := common.NewLocation(randX, randY, randZ)

			var randLoc2 common.Location

			switch rand.Intn(6) {
			case 0:
				randLoc2 = common.NewLocation(randX+1, randY, randZ)
			case 1:
				randLoc2 = common.NewLocation(randX-1, randY, randZ)
			case 2:
				randLoc2 = common.NewLocation(randX, randY+1, randZ)
			case 3:
				randLoc2 = common.NewLocation(randX, randY-1, randZ)
			case 4:
				randLoc2 = common.NewLocation(randX, randY, randZ+1)
			case 5:
				randLoc2 = common.NewLocation(randX, randY, randZ-1)
			}

			lab.Connect(randLoc, randLoc2)
		}

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
