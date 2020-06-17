package display

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gotk3/gotk3/glib"

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
	Window         *gtk.Window
	Visualizer     LabyrinthVisualizer
	Generator      generator.LabGenerator
	generatorSteps []common.Pair
	SolverFunc     solverFunction
	SolvePath      []common.Location
	solverSteps    []common.Pair

	glArea                                                *gtk.GLArea
	startTime                                             time.Time
	cameraPosition, lookAtCenter, upVector, lightPosition mgl32.Vec3
	lab                                                   common.Labyrinth
	projectionMatrix                                      mgl32.Mat4
	constructor                                           CubeConstructor
	transform                                             Transform
	rotateAxisX, rotateAxisY                              mgl32.Vec3
	draggingEnabled                                       bool
	shouldStep                                            bool
}

const (
	fov        float32 = 45
	nearCutoff float32 = 0.1
	farCutoff  float32 = 100
	jumpThresh float64 = 100
	stepTime           = 100
)

func generators() map[string]generator.LabGenerator {
	return map[string]generator.LabGenerator{
		"DepthFirst":   generator.NewDepthFirstGenerator(),
		"BreadthFirst": generator.NewBreadthFirstGenerator(),
	}
}

func solvers() map[string]solverFunction {
	return map[string]solverFunction{
		"Recursive": func(a common.Labyrinth, b, c common.Location) []common.Location {
			return solver.RecursiveSolver(a, b, c, false)
		},
	}
}

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
		rotateAxisX: mgl32.Vec3{
			0, 1, 0,
		},
		rotateAxisY: mgl32.Vec3{
			1, 0, -1,
		}.Normalize(),
		Window:    &win.Window,
		lab:       nil,
		Generator: generator.NewDepthFirstGenerator(),
		SolverFunc: func(labyrinth common.Labyrinth, from, to common.Location) []common.Location {
			return solver.RecursiveSolver(labyrinth, from, to, false)
		},
		transform: TransformIdent(),
	}

	signals := map[string]interface{}{
		"gl_init":                      wnd.realize,   // Called on Window Creation
		"gl_draw":                      wnd.render,    // Window Redraw
		"gl_fini":                      wnd.unrealize, // Window Deletion
		"on_generate_random_labyrinth": wnd.generateRandomLab,
		"on_switch_dragging":           wnd.switchBetweenDraggingAndAutoRotate,
		"on_show_group_changed":        wnd.onShowGroupChanged,
	}

	builder.ConnectSignals(signals)
	initChooseSubmenu(builder, &wnd)
	initDraggingFunctionality(glArea, &wnd)

	return &wnd
}

func initChooseSubmenu(builder *gtk.Builder, wnd *MainWindow) {
	solverMenuObj, err := builder.GetObject("choose_solver_menu")
	FatalIfError("choose solver submenu not found in glarea.ui", err)
	generatorMenuObj, err := builder.GetObject("choose_generator_menu")
	FatalIfError("choose generator submenu not found in glarea.ui", err)

	var solverMenu *gtk.MenuItem

	var generatorMenu *gtk.MenuItem

	if result, ok := solverMenuObj.(*gtk.MenuItem); ok {
		solverMenu = result
	} else {
		log.Fatal("\"choose_solver_menu\" is not a *MenuItem")
	}

	if result, ok := generatorMenuObj.(*gtk.MenuItem); ok {
		generatorMenu = result
	} else {
		log.Fatal("\"choose_generator_menu\" is not a *MenuItem")
	}

	i := 0
	names := make([]string, len(solvers()))

	for name := range solvers() {
		names[i] = name
		i++
	}

	solverSubMenu := createMenuItems(names, wnd.onSolverChanged)

	i = 0
	names = make([]string, len(generators()))

	for name := range generators() {
		names[i] = name
		i++
	}

	generatorSubMenu := createMenuItems(names, wnd.onGeneratorChanged)

	solverMenu.SetSubmenu(solverSubMenu)
	generatorMenu.SetSubmenu(generatorSubMenu)
}

func createMenuItems(names []string, fun func(*gtk.RadioMenuItem)) *gtk.Menu {
	var baseItem *gtk.RadioMenuItem = nil

	menu, err := gtk.MenuNew()

	FatalIfError("Could not create new Menu", err)

	for _, name := range names {
		var item *gtk.RadioMenuItem

		if baseItem == nil {
			backingList := glib.SList{}
			item, err = gtk.RadioMenuItemNewWithLabel(&backingList, name)
			baseItem = item
		} else {
			item, err = gtk.RadioMenuItemNewWithLabelFromWidget(baseItem, name)
		}

		FatalIfError("Could not create menu entry for "+name, err)

		_, err = item.Connect("activate", fun)
		FatalIfError("Could not connect to activate on menu item "+name, err)

		menu.Append(item)
	}

	return menu
}

func (wnd *MainWindow) onGeneratorChanged(item *gtk.RadioMenuItem) {
	if !item.GetActive() {
		return
	}

	log.Println("Generator changing to " + item.GetLabel())
}

func (wnd *MainWindow) onSolverChanged(item *gtk.RadioMenuItem) {
	if !item.GetActive() {
		return
	}

	log.Println("Generator changing to " + item.GetLabel())
}

func (wnd *MainWindow) onShowGroupChanged(item *gtk.RadioMenuItem) {
	if !item.GetActive() {
		return
	}

	name, err := item.GetName()

	FatalIfError("Selected Menu Item has no name!", err)
	log.Printf("Show Group Changed: %v", item.GetLabel())

	activateStepCallback := func() {
		wnd.shouldStep = true
		_, err := glib.TimeoutAdd(stepTime, wnd.stepCallback)
		FatalIfError("Could not install step callback", err)
	}

	switch name {
	case "show_solver_path":
		{
			wnd.shouldStep = false
			wnd.Visualizer.SetSteps(nil, nil)
			wnd.Visualizer.SetPath(wnd.SolvePath)
			break
		}
	case "show_solver_algorithm":
		{
			wnd.Visualizer.SetPath(nil)
			wnd.Visualizer.SetSteps(wnd.solverSteps, nil) //TODO
			activateStepCallback()
			break
		}
	case "show_generator_algorithm":
		{
			wnd.Visualizer.SetPath(nil)
			wnd.Visualizer.SetSteps(wnd.generatorSteps, GeneratorColorConverter())
			activateStepCallback()
			break
		}
	default:
		{
			log.Println("Show not recognized: " + name)
		}
	}
}

func (wnd *MainWindow) switchBetweenDraggingAndAutoRotate() {
	wnd.draggingEnabled = !wnd.draggingEnabled
}

func initDraggingFunctionality(glArea *gtk.GLArea, wnd *MainWindow) {
	glArea.AddEvents(int(gdk.POINTER_MOTION_MASK) | int(gdk.BUTTON_PRESS_MASK) | int(gdk.BUTTON_RELEASE_MASK))

	var dragging bool

	lastX := 0.0
	lastY := 0.0

	_, err := glArea.Connect("button_press_event", func() {
		dragging = true
	})

	FatalIfError("Could not connect to button_press_event signal", err)

	_, err = glArea.Connect("button_release_event", func() {
		dragging = false
	})

	FatalIfError("Could not connect to button_release_event signal", err)

	_, err = glArea.Connect("motion_notify_event", func(widget *gtk.GLArea, event *gdk.Event) {
		if wnd.draggingEnabled && dragging {
			motionEvent := gdk.EventMotionNewFromEvent(event)
			x, y := motionEvent.MotionVal()
			dX := x - lastX
			dY := y - lastY

			lastX = x
			lastY = y
			if math.Abs(dX) < jumpThresh && math.Abs(dY) < jumpThresh {
				wnd.mouseDrag(dX, dY)
			}
		}
	})

	FatalIfError("Could not connect to motion_notify_event signal", err)
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

func (wnd *MainWindow) mouseDrag(x, y float64) {
	normX := x / float64(wnd.glArea.GetAllocatedWidth())
	normY := y / float64(wnd.glArea.GetAllocatedHeight())

	rotX := mgl32.QuatRotate(float32(normX), wnd.rotateAxisX)
	rotY := mgl32.QuatRotate(float32(normY), wnd.rotateAxisY)
	rot := rotY.Mul(rotX)

	wnd.transform.rotation = rot.Mul(wnd.transform.rotation)
}

func (wnd *MainWindow) SetLabyrinth(lab *common.Labyrinth) {
	if lab == nil {
		return
	}

	wnd.lab = *lab

	wnd.shouldStep = false
	wnd.Visualizer = NewLabyrinthVisualizer(&wnd.lab, wnd.constructor)

	labMaxX, labMaxY, labMaxZ := wnd.lab.GetMaxLocation().As3DCoordinates()

	labSizeX := float32(labMaxX + 1)
	labSizeY := float32(labMaxY + 1)
	labSizeZ := float32(labMaxZ + 1)

	labCenter := mgl32.Vec3{float32(labMaxX) / 2.0, float32(labMaxY) / 2.0, float32(labMaxZ) / 2.0}

	wnd.transform.SetTranslation(-labCenter.X(), -labCenter.Y(), -labCenter.Z())
	wnd.transform.rotation = mgl32.QuatIdent()

	wnd.lightPosition = mgl32.Vec3{
		-labSizeX, labSizeY, labSizeZ,
	}

	wnd.cameraPosition = mgl32.Vec3{
		labSizeX, labSizeY, labSizeZ,
	}

	from := common.NewLocation(0, 0, 0)
	to := common.NewLocation(labMaxX, labMaxY, labMaxZ)

	wnd.SolvePath = wnd.SolverFunc(wnd.lab, from, to)
}

// Called by gtk every time the window has to draw its contents.
func (wnd *MainWindow) render() {
	if !wnd.Visualizer.IsValid() {
		return
	}

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	view := mgl32.LookAtV(wnd.cameraPosition, wnd.lookAtCenter, wnd.upVector)

	transform := wnd.transform.AsMatrix()

	for _, cube := range wnd.Visualizer.cubes {
		cube.draw(&view, &wnd.projectionMatrix, &transform, wnd.lightPosition)
	}
}

func (wnd *MainWindow) generateRandomLab() {
	const startLabSize = 10

	rand.Seed(time.Now().UnixNano())

	randInt := func(max uint) uint { return uint(rand.Intn(int(max))) }
	furthestPoint := common.NewLocation(randInt(startLabSize), randInt(startLabSize), randInt(startLabSize))

	lab, steps := generator.NewDepthFirstGenerator().GenerateLabyrinth(furthestPoint)
	wnd.generatorSteps = steps

	wnd.SetLabyrinth(&lab)
}

// Called by gtk 60 times per second (once per "tick")
func (wnd *MainWindow) update(widget *gtk.Widget, clock *gdk.FrameClock, _ uintptr) bool {
	if !wnd.draggingEnabled {
		angle := float32(clock.GetFrameTime()) / 1000000.0 //nolint:gomnd
		wnd.transform.SetRotation(angle, wnd.upVector)
	}

	widget.QueueDraw()

	return true
}

func (wnd *MainWindow) stepCallback() bool {
	if wnd.shouldStep && wnd.Visualizer.IsValid() {
		wnd.Visualizer.DoStep()
	}

	return wnd.shouldStep
}

// Called on window destruction
func (wnd *MainWindow) unrealize() {
	log.Println("Unrealizing Main Window")
}
