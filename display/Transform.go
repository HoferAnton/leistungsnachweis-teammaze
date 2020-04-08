package display

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Transform struct {
	translation, scale mgl32.Mat4
	rotation           mgl32.Quat
}

func (t *Transform) AsMatrix() mgl32.Mat4 {
	return t.translation.Mul4(t.rotation.Mat4().Mul4(t.scale))
}

func (t *Transform) SetRotation(rotation mgl32.Mat4) {
	t.rotation = mgl32.Mat4ToQuat(rotation)
}

func (t *Transform) SetScale(scale mgl32.Mat4) {
	t.scale = scale
}

func (t *Transform) SetTranslation(translation mgl32.Mat4) {
	t.translation = translation
}
