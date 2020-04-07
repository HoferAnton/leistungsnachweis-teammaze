package display

import (
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"log"
	"math/rand"
	"time"

	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type MainWindow struct {
	glArea                                 *gtk.GLArea
	startTime                              time.Time
	cube, cube2                            Cube
	cameraPosition, lookAtCenter, upVector mgl32.Vec3
	Window                                 *gtk.Window
	Visualizer                             LabyrinthVisualizer
	projectionMatrix                       mgl32.Mat4
}

const fov float32 = 45
const nearCutoff = 0.1
const farCutoff = 100
const labyrinthSize = 5
const numConnections = 100

func CreateMainWindow() MainWindow {
	builder, err := gtk.BuilderNewFromFile("display/ui/glarea.ui")

	FatalIfError("Could not create GTK Builder: ", err)

	obj, err := builder.GetObject("main_window")
	FatalIfError("Could not get main_window: ", err)

	win, err := asWindow(obj)
	FatalIfError("Could not use main_window: ", err)

	glAreaObject, err := builder.GetObject("gl_drawing_area")
	FatalIfError("Could not find gl_drawing_area: ", err)

	var glArea *gtk.GLArea

	var ok bool
	if glArea, ok = glAreaObject.(*gtk.GLArea); !ok {
		log.Fatal("gl_drawing_area is not a GLArea")
	}

	wnd := MainWindow{
		glArea: glArea,
		cameraPosition: mgl32.Vec3{
			labyrinthSize, labyrinthSize, labyrinthSize,
		},
		lookAtCenter: mgl32.Vec3{
			0, 0, 0,
		},
		upVector: mgl32.Vec3{
			0, 1, 0,
		},
		Window: &win.Window,
	}

	signals := map[string]interface{}{
		"gl_init": wnd.realize,   // Called on Window Creation
		"gl_draw": wnd.render,    // Window Redraw
		"gl_fini": wnd.unrealize, // Window Deletion
	}

	builder.ConnectSignals(signals)

	return wnd
}

func (wnd *MainWindow) realize() {
	log.Println("Realizing Main Window")

	wnd.startTime = time.Now()

	wnd.glArea.SetHasDepthBuffer(true)

	wnd.glArea.MakeCurrent()

	wnd.glArea.AddTickCallback(wnd.update, uintptr(0))

	err := gl.Init()

	FatalIfError("Could not init OpenGL: ", err)

	aspectRatio := float32(wnd.glArea.GetAllocatedWidth()) / float32(wnd.glArea.GetAllocatedHeight())
	wnd.projectionMatrix = mgl32.Perspective(fov, aspectRatio, nearCutoff, farCutoff)

	lab := common.NewLabyrinth(common.NewLocation(labyrinthSize-1, labyrinthSize-1, labyrinthSize-1))

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

	wnd.Visualizer = NewLabyrinthVisualizer(&lab)
	for _, cube := range wnd.Visualizer.cubes {
		log.Printf("%v\n", cube)
	}
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0, 0, 0, 1)
}

func (wnd *MainWindow) render() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	view := mgl32.LookAtV(wnd.cameraPosition, wnd.lookAtCenter, wnd.upVector)
	labCenter := mgl32.Vec3{(labyrinthSize - 1) / 2.0, (labyrinthSize - 1) / 2.0, (labyrinthSize - 1) / 2.0}

	for _, cube := range wnd.Visualizer.cubes {
		cube.draw(view, wnd.projectionMatrix, labCenter, wnd.cameraPosition, time.Since(wnd.startTime))
	}
}

func (wnd *MainWindow) update(widget *gtk.Widget, _ *gdk.FrameClock, _ uintptr) bool {
	widget.QueueDraw()
	return true
}

func (wnd *MainWindow) unrealize() {
	log.Println("Unrealizing Main Window")
}
