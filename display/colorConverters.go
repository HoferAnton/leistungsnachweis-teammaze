package display

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/generator"
)

func GeneratorColorConverter() StepColorConverter {
	return NewColorConverter(map[string]mgl32.Vec4{
		generator.Start:     {0, 0.33, 1, 1},
		generator.Discover:  {0, 1, 1, 1},
		generator.Backtrack: {0.5, 0, 0, 1},
		generator.Select:    {0, 0.5, 0, 1},
		generator.Add:       {1, 0.33, 0, 1},
	})
}
