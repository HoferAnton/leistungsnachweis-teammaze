package display

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Transform struct {
	translation, scale mgl32.Mat4
	rotation           mgl32.Quat
}

func TransformIdent() Transform {
	return Transform{
		translation: mgl32.Ident4(),
		scale:       mgl32.Ident4(),
		rotation:    mgl32.QuatIdent(),
	}
}

func (t *Transform) AsMatrix() mgl32.Mat4 {
	return t.translation.Mul4(t.rotation.Mat4().Mul4(t.scale))
}

func (t *Transform) GetTranslation() mgl32.Vec3 {
	return t.translation.Mul4x1(mgl32.Vec4{0, 0, 0, 1}).Vec3()
}

func (t *Transform) SetRotation(angle float32, axis mgl32.Vec3) {
	t.rotation = mgl32.QuatRotate(angle, axis)
}

func (t *Transform) SetScale(x, y, z float32) {
	t.scale = mgl32.Scale3D(x, y, z)
}

func (t *Transform) SetTranslation(x, y, z float32) {
	t.translation = mgl32.Translate3D(x, y, z)
}
