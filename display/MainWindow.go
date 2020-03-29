package display

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

type MainWindow struct {
}

func CreateMainWindow(width int, height int) MainWindow {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

	if err != nil {
		log.Fatal("Unable to create window: ", err)
	}

	win.SetTitle("Example")
	_, err = win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	if err != nil {
		log.Fatal("Could not connect to destroy signal on main window: ", err)
	}

	win.SetDefaultSize(width, height)
	win.ShowAll()
	gtk.Main()

	return MainWindow{}
}
