package solver

import (
	"reflect"
	"testing"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/generator"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/printer"
)

func TestConcurrentSolver(t *testing.T) {
	lab, _ := generator.NewDepthFirstGenerator().GenerateLabyrinth(common.NewLocation(uint(10), uint(7), uint(0)))
	rs := RecursiveSolver(lab, common.NewLocation(0, 0, 0), lab.GetMaxLocation(), false)
	cs := ConcurrentSolver(lab, common.NewLocation(0, 0, 0), lab.GetMaxLocation(), false)
	if !reflect.DeepEqual(rs, cs) {
		srs, _ := printer.Print2D(lab, rs)
		scs, _ := printer.Print2D(lab, cs)
		t.Errorf("rs: \n%v\ncs: \n%v", srs, scs)
	}
}
