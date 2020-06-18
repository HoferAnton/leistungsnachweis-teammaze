package display

import (
	"testing"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/generator"
)

func TestGeneratorColorConverter(t *testing.T) {
	lab, steps := generator.NewDepthFirstGenerator().GenerateLabyrinth(common.NewLocation(5, 5, 5))
	vis := NewLabyrinthVisualizer(&lab, testingCubeConstructor)
	sut := GeneratorColorConverter()

	for _, step := range steps {
		if _, color := sut.StepToColor(step, &vis); color == (mgl32.Vec4{}) {
			t.Errorf("got null color!")
		}
	}

	lab, steps = generator.NewBreadthFirstGenerator().GenerateLabyrinth(common.NewLocation(5, 5, 5))
	vis = NewLabyrinthVisualizer(&lab, testingCubeConstructor)

	for _, step := range steps {
		if _, color := sut.StepToColor(step, &vis); color == (mgl32.Vec4{}) {
			t.Errorf("got null color!")
		}
	}
}
