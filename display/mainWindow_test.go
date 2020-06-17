package display

import (
	"testing"

	"github.com/gotk3/gotk3/gtk"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/solver"
)

func TestConsts(t *testing.T) {
	if len(generators()) == 0 {
		t.Error("")
	}

	for name, gen := range generators() {
		if name == "" || gen == nil {
			t.Error("")
		}
	}

	if len(solvers()) == 0 {
		t.Error("")
	}

	for name, item := range solvers() {
		if name == "" || item == nil {
			t.Error("")
		}
	}

	wnd := testWnd()

	if len(signals(&wnd)) == 0 {
		t.Error("")
	}

	for name, fun := range signals(&wnd) {
		if name == "" || fun == nil {
			t.Error("")
		}
	}
}

func testWnd() MainWindow {
	return MainWindow{
		constructor: testingCubeConstructor,
		SolverFunc:  solver.RecursiveSolverSteps,
	}
}

func TestCreateMainWindow(t *testing.T) {
	gtk.Init(nil)

	wnd := CreateMainWindow("ui/glarea.ui")

	if wnd.glArea == nil || wnd.Window == nil || wnd.labelContainer == nil ||
		wnd.Generator == nil || wnd.SolverFunc == nil {
		t.Error("")
	}
}

func TestMainWindow_generateRandomLab(t *testing.T) {
	wnd := testWnd()

	wnd.generateRandomLab()

	if wnd.lab.GetMaxLocation() == common.NewLocation(0, 0, 0) {
		t.Error("")
	}

	if wnd.generatorSteps == nil || len(wnd.generatorSteps) == 0 {
		t.Error("")
	}

	if !wnd.Visualizer.IsValid() {
		t.Error("")
	}

	if wnd.solverSteps == nil || len(wnd.solverSteps) == 0 {
		t.Error("")
	}
}
