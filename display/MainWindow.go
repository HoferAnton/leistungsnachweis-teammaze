package display

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

type MainWindow struct {
}

func CreateMainWindow(width int, height int) MainWindow {
	gtk.Init(nil)
	doOnMain(func() {
		err := gl.Init()

		FatalIfError("Could not initialize OpenGL: ", err)
	})

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

	FatalIfError("Unable to create window: ", err)

	win.SetTitle("Example")
	_, err = win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	FatalIfError("Could not connect to destroy signal on main window: ", err)

	win.SetDefaultSize(width, height)
	win.ShowAll()
	gtk.Main()

	return MainWindow{}
}

func Realize() {
	log.Println("Realizing Main Window")
}

func Render() {
	log.Println("Rendering Main Window")
}

func Unrealize() {
	log.Println("Unrealizing Main Window")
}
