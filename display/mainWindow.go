package display

import (
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"log"
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
			8, 6, 8,
		},
		lookAtCenter: mgl32.Vec3{
			0, 2, 0,
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

const fov float32 = 45
const nearCutoff = 0
const farCutoff = 100
const labyrinthSize = 2

func (wnd *MainWindow) realize() {
	log.Println("Realizing Main Window")

	wnd.startTime = time.Now()

	wnd.glArea.MakeCurrent()

	wnd.glArea.AddTickCallback(wnd.update, uintptr(0))

	err := gl.Init()

	FatalIfError("Could not init OpenGL: ", err)

	aspectRatio := float32(wnd.glArea.GetAllocatedWidth()) / float32(wnd.glArea.GetAllocatedHeight())
	wnd.projectionMatrix = mgl32.Perspective(fov, aspectRatio, nearCutoff, farCutoff)

	lab := common.NewLabyrinth(common.NewLocation(labyrinthSize-1, labyrinthSize-1, labyrinthSize-1))

	lab.Connect(common.NewLocation(0, 0, 0), common.NewLocation(0, 1, 0))
	lab.Connect(common.NewLocation(0, 1, 0), common.NewLocation(1, 1, 0))
	lab.Connect(common.NewLocation(1, 1, 0), common.NewLocation(1, 1, 1))
	lab.Connect(common.NewLocation(1, 1, 1), common.NewLocation(0, 1, 1))

	//rand.Seed(time.Now().UnixNano())

	//for i := 0; i < 25;i++ {
	//	loc1 := common.NewLocation(uint(rand.Intn(labyrinthSize)), uint(rand.Intn(labyrinthSize)), uint(rand.Intn(labyrinthSize)))
	//	loc2 := common.NewLocation(uint(rand.Intn(labyrinthSize)), uint(rand.Intn(labyrinthSize)), uint(rand.Intn(labyrinthSize)))
	//
	//	if loc1.Compare(loc2) {
	//		continue
	//	}
	//
	//	lab.Connect(loc1, loc2)
	//}

	wnd.Visualizer = NewLabyrinthVisualizer(&lab)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1, 1, 1, 1)
}

func (wnd *MainWindow) render() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	view := mgl32.LookAtV(wnd.cameraPosition, wnd.lookAtCenter, wnd.upVector)

	for _, cube := range wnd.Visualizer.cubes {
		cube.draw(view, wnd.projectionMatrix, time.Since(wnd.startTime))
	}
}

func (wnd *MainWindow) update(widget *gtk.Widget, _ *gdk.FrameClock, _ uintptr) bool {
	widget.QueueDraw()
	return true
}

func (wnd *MainWindow) unrealize() {
	log.Println("Unrealizing Main Window")
}
