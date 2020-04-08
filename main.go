package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/display"
)

const appID = "com.github.ob-algdatii-20ss.leistungsnachweis-teammaze"

const labyrinthSize = 5
const numConnections = 100

type directionEnum int

const (
	right    directionEnum = 0
	left     directionEnum = 1
	up       directionEnum = 2
	down     directionEnum = 3
	backward directionEnum = 4
	forward  directionEnum = 5
)

func main() {
	runtime.LockOSThread()

	log.Println("Execution path: ", os.Args[0])
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

			randLoc2 := getRandomNeighboringLocation(randX, randY, randZ)

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

func getRandomNeighboringLocation(randX uint, randY uint, randZ uint) common.Location {
	switch directionEnum(rand.Intn(6)) { //nolint:gomnd Number of directions in three dimensions is always 6
	case right:
		return common.NewLocation(randX+1, randY, randZ)
	case left:
		return common.NewLocation(randX-1, randY, randZ)
	case up:
		return common.NewLocation(randX, randY+1, randZ)
	case down:
		return common.NewLocation(randX, randY-1, randZ)
	case backward:
		return common.NewLocation(randX, randY, randZ+1)
	case forward:
		return common.NewLocation(randX, randY, randZ-1)
	default:
		panic("Rolled an unknown direction")
	}
}
