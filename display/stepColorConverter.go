package display

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

type StepColorConverter interface {
	StepToColor(common.Pair, *LabyrinthVisualizer) (*Cube, mgl32.Vec4)
	ColorMap() map[string]mgl32.Vec4
}

type MappingColorConverter struct {
	mapping map[string]mgl32.Vec4
}

func NewColorConverter(mapping map[string]mgl32.Vec4) MappingColorConverter {
	return MappingColorConverter{
		mapping: mapping,
	}
}

func (t MappingColorConverter) StepToColor(step common.Pair, vis *LabyrinthVisualizer) (*Cube, mgl32.Vec4) {
	if step.GetFirst() == nil || step.GetSecond() == nil || vis == nil {
		panic("nil not allowed in arguments")
	}

	location := step.GetFirst().(common.Location)
	symbol := step.GetSecond().(string)

	x, y, z := location.As3DCoordinates()
	location3D := mgl32.Vec3{float32(x), float32(y), float32(z)}

	return vis.GetCubeAt(location3D), t.mapping[symbol]
}

func (t MappingColorConverter) ColorMap() map[string]mgl32.Vec4 {
	return t.mapping
}
