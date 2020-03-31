package display

import (
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/gotk3/gotk3/gtk"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"log"
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

var cube Cube
var glArea *gtk.GLArea

func realize() {
	log.Println("Realizing Main Window")

	glArea.MakeCurrent()

	err := gl.Init()

	FatalIfError("Could not init OpenGL: ", err)

	_, err = CreateProgram("display/shaders/simple_vertex.glsl", "display/shaders/simple_fragment.glsl")
	FatalIfError("Failed to compile shader program: ", err)

	cube = NewCube(common.NewLocation(1, 1, 1), 1, 1, 1)
}

func render() {
	log.Println("Rendering Main Window")

	cube.Draw()
}

func unrealize() {
	log.Println("Unrealizing Main Window")
}
