package display

import (
	"log"
	"time"

	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

func CreateMainWindow() *gtk.Window {
	builder, err := gtk.BuilderNewFromFile("display/ui/glarea.ui")

	FatalIfError("Could not create GTK Builder: ", err)

	signals := map[string]interface{}{
		"gl_init": realize,   // Called on Window Creation
		"gl_draw": render,    // Window Redraw
		"gl_fini": unrealize, // Window Deletion
	}

	builder.ConnectSignals(signals)

	obj, err := builder.GetObject("main_window")

	FatalIfError("Could not get main_window: ", err)

	win, err := asWindow(obj)

	FatalIfError("Could not use main_window: ", err)

	glAreaObject, err := builder.GetObject("gl_drawing_area")

	FatalIfError("Could not find gl_drawing_area: ", err)

	var ok bool

	if glArea, ok = glAreaObject.(*gtk.GLArea); !ok {
		log.Fatal("gl_drawing_area is not a GLArea")
	}

	return &(win.Window)
}

var glArea *gtk.GLArea

var startTime time.Time

const fov float32 = 45
const nearCutoff = 0
const farCutoff = 100

func realize() {
	log.Println("Realizing Main Window")
	startTime = time.Now()

	glArea.MakeCurrent()

	glArea.AddTickCallback(update, uintptr(0))

	err := gl.Init()

	FatalIfError("Could not init OpenGL: ", err)

	aspectRatio := float32(glArea.GetAllocatedWidth()) / float32(glArea.GetAllocatedHeight())

	projection := mgl32.Perspective(fov, aspectRatio, nearCutoff, farCutoff)
	cube = NewCube(common.NewLocation(0, 0, 0), 1, 1, 1, projection)
	cube2 = NewCube(common.NewLocation(1, 1, 1), 0.5, 0.5, 0.5, projection)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1, 1, 1, 1)
}

var cube Cube
var cube2 Cube

var cameraPosition = mgl32.Vec3{
	4, 3, 4,
}

var worldCenter = mgl32.Vec3{
	0, 0, 0,
}

var upVector = mgl32.Vec3{
	0, 1, 0,
}

func render() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	view := mgl32.LookAtV(cameraPosition, worldCenter, upVector)
	cube.draw(view, time.Since(startTime))
	cube2.draw(view, time.Since(startTime))
}

func update(widget *gtk.Widget, _ *gdk.FrameClock, _ uintptr) bool {
	widget.QueueDraw()
	return true
}

func unrealize() {
	log.Println("Unrealizing Main Window")
}
