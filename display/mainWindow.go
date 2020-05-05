package display

import (
	"log"
	"math/rand"
	"time"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/solver"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/generator"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"

	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// This probably belongs in the solver interface
type solverFunction = func(common.Labyrinth, common.Location, common.Location) []common.Location

type MainWindow struct {
	glArea                                                *gtk.GLArea
	startTime                                             time.Time
	cameraPosition, lookAtCenter, upVector, lightPosition mgl32.Vec3
	Window                                                *gtk.Window
	lab                                                   common.Labyrinth
	Visualizer                                            LabyrinthVisualizer
	projectionMatrix                                      mgl32.Mat4
	Generator                                             generator.LabGenerator
	SolverFunc                                            solverFunction
	SolvePath                                             []common.Location
	constructor                                           CubeConstructor
}

const fov float32 = 45
const nearCutoff = 0.1
const farCutoff = 100

// MainWindow constructor:
// Loads ui configuration from ui/glarea.ui (gtk xml file / edit per hand or with glade)
// Initializes OpenGL with GLContext from GLArea
// Connects GTK Signals to Callback functions

func CreateMainWindow() *MainWindow {
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
		lookAtCenter: mgl32.Vec3{
			0, 0, 0,
		},
		upVector: mgl32.Vec3{
			0, 1, 0,
		},
		Window:    &win.Window,
		lab:       nil,
		Generator: generator.NewDepthFirstGenerator(),
		SolverFunc: func(labyrinth common.Labyrinth, from, to common.Location) []common.Location {
			return solver.RecursiveSolver(labyrinth, from, to, false)
		},
	}

	signals := map[string]interface{}{
		"gl_init":                      wnd.realize,   // Called on Window Creation
		"gl_draw":                      wnd.render,    // Window Redraw
		"gl_fini":                      wnd.unrealize, // Window Deletion
		"on_generate_random_labyrinth": wnd.generateRandomLab,
	}

	builder.ConnectSignals(signals)

	return &wnd
}

// Called before the window is shown.
func (wnd *MainWindow) realize() {
	log.Println("Realizing Main Window")

	wnd.startTime = time.Now()
	wnd.glArea.MakeCurrent()
	wnd.glArea.AddTickCallback(wnd.update, uintptr(0))

	err := gl.Init()

	FatalIfError("Could not init OpenGL: ", err)

	aspectRatio := float32(wnd.glArea.GetAllocatedWidth()) / float32(wnd.glArea.GetAllocatedHeight())
	wnd.projectionMatrix = mgl32.Perspective(fov, aspectRatio, nearCutoff, farCutoff)

	shaderProgram, err := CreateProgram("display/shaders/simple_vertex.glsl", "display/shaders/simple_fragment.glsl")

	gl.BindFragDataLocation(shaderProgram, 0, gl.Str("colorOut\x00"))
	FatalIfError("Could not create shader program", err)

	wnd.constructor = GetCubeConstructor(shaderProgram)

	wnd.generateRandomLab()

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0, 0, 0, 1)
}

func (wnd *MainWindow) SetLabyrinth(lab *common.Labyrinth) {
	if lab == nil {
		return
	}

	wnd.lab = *lab
	wnd.Visualizer = NewLabyrinthVisualizer(&wnd.lab, wnd.constructor)

	labMaxX, labMaxY, labMaxZ := wnd.lab.GetMaxLocation().As3DCoordinates()

	labSizeX := float32(labMaxX + 1)
	labSizeY := float32(labMaxY + 1)
	labSizeZ := float32(labMaxZ + 1)

	wnd.lightPosition = mgl32.Vec3{
		-labSizeX, labSizeY, labSizeZ,
	}

	wnd.cameraPosition = mgl32.Vec3{
		labSizeX, labSizeY, labSizeZ,
	}

	from := common.NewLocation(0, 0, 0)
	to := common.NewLocation(labMaxX, labMaxY, labMaxZ)

	wnd.Visualizer.SetPath(wnd.SolverFunc(wnd.lab, from, to))
}

// Called by gtk every time the window has to draw its contents.
func (wnd *MainWindow) render() {
	if !wnd.Visualizer.IsValid() {
		return
	}

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	labMaxX, labMaxY, labMaxZ := wnd.lab.GetMaxLocation().As3DCoordinates()

	view := mgl32.LookAtV(wnd.cameraPosition, wnd.lookAtCenter, wnd.upVector)
	labCenter := mgl32.Vec3{float32(labMaxX) / 2.0, float32(labMaxY) / 2.0, float32(labMaxZ) / 2.0}

	transform := mgl32.HomogRotate3DY(float32(time.Since(wnd.startTime).Seconds())).
		Mul4(mgl32.Translate3D(-labCenter.X(), -labCenter.Y(), -labCenter.Z()))

	for _, cube := range wnd.Visualizer.cubes {
		cube.draw(&view, &wnd.projectionMatrix, &transform, wnd.lightPosition)
	}
}

func (wnd *MainWindow) generateRandomLab() {
	const startLabSize = 20

	rand.Seed(time.Now().UnixNano())

	randInt := func(max uint) uint { return uint(rand.Intn(int(max))) }
	furthestPoint := common.NewLocation(randInt(startLabSize), randInt(startLabSize), randInt(startLabSize))

	lab, _ := generator.NewDepthFirstGenerator().GenerateLabyrinth(furthestPoint)
	lab.Connect(common.NewLocation(0, 0, 0), common.NewLocation(0, 0, 1))
	wnd.SetLabyrinth(&lab)
}

// Called by gtk 60 times per second (once per "tick")
func (wnd *MainWindow) update(widget *gtk.Widget, _ *gdk.FrameClock, _ uintptr) bool {
	widget.QueueDraw()
	return true
}

// Called on window destruction
func (wnd *MainWindow) unrealize() {
	log.Println("Unrealizing Main Window")
}
