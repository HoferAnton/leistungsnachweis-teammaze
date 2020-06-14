package display

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

type StepColorConverter interface {
	StepToColor(common.Pair, []Cube) (*Cube, mgl32.Vec4)
	StepColorMap() map[string]mgl32.Vec4
}
