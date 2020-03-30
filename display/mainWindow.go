package display

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"log"
	"unsafe"

	"github.com/gotk3/gotk3/gtk"
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

	return &(win.Window)
}

var cube Cube

func realize() {
	log.Println("Realizing Main Window")
	err := gl.Init()

	gl.DebugMessageCallback(func(source uint32, gltype uint32, id uint32, severity uint32, length int32, message string, userParam unsafe.Pointer) {
		log.Println("Received Debug Message: ", message)
	}, nil)

	FatalIfError("Could not initialize OpenGL: ", err)
	cube = NewCube(common.NewLocation(1, 1, 1), 1, 1, 1)
}

func render() {
	log.Println("Rendering Main Window")
}

func unrealize() {
	log.Println("Unrealizing Main Window")
}
