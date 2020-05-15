package display

import (
	"math/rand"
	"testing"
	"time"

	"github.com/go-gl/mathgl/mgl32"
)

func TestTransform_SetRotation(t *testing.T) {
	sut := Transform{}

	rand.Seed(time.Now().UnixNano())

	angle := rand.Float32()
	axis := randomVec3()

	sut.SetRotation(angle, axis)

	compareMatrices(t, mgl32.QuatRotate(angle, axis).Mat4(), sut.rotation.Mat4(), "Rotation not set correctly")
}

func TestTransform_SetScale(t *testing.T) {
	sut := Transform{}

	rand.Seed(time.Now().UnixNano())

	scale := randomVec3()

	sut.SetScale(scale.Elem())

	compareMatrices(t, mgl32.Scale3D(scale.Elem()), sut.scale, "Scale not set correctly")
}

func TestTransform_SetTranslation(t *testing.T) {
	sut := Transform{}

	rand.Seed(time.Now().UnixNano())

	translation := randomVec3()

	sut.SetTranslation(translation.Elem())

	compareMatrices(t, mgl32.Translate3D(translation.Elem()), sut.translation, "Translation not set correctly")
}

func TestTransform_AsMatrix(t *testing.T) {
	sut := Transform{}

	rand.Seed(time.Now().UnixNano())

	scale := randomVec3()
	translate := randomVec3()
	axis := randomVec3()
	angle := rand.Float32()

	sut.SetTranslation(translate.Elem())
	sut.SetScale(scale.Elem())
	sut.SetRotation(angle, axis)

	want := mgl32.QuatRotate(angle, axis).Mat4().Mul4( // Rotation *
		mgl32.Translate3D(translate.Elem()).Mul4( // Translation *
			mgl32.Scale3D(scale.Elem()))) // Scale (Order is important, else your transformations get mixed up)
	got := sut.AsMatrix()

	compareMatrices(t, want, got, "Transform matrix does not match")
}

func TestTransformIdent(t *testing.T) {
	sut := TransformIdent()

	testVec := randomVec3().Vec4(1) // w coordinate is always 1 in shaders.

	got := sut.AsMatrix().Mul4x1(testVec)

	if got != testVec {
		t.Errorf("Identity Transform changed vector:\ngot: %v\nexpected: %v", got, testVec)
	}
}

func TestTransform_GetTranslation(t *testing.T) {
	sut := TransformIdent()

	x, y, z := randomVec3().Elem()

	sut.SetTranslation(x, y, z)

	got := sut.GetTranslation()
	want := mgl32.Vec3{x, y, z}

	if got != want {
		t.Errorf("Identity Transform changed vector:\ngot: %v\nexpected: %v", got, want)
	}
}

//Helpers:

func compareMatrices(t *testing.T, want, got mgl32.Mat4, message string) {
	if !got.ApproxEqual(want) {
		t.Errorf("%s:\ngot:\n%v\nexpected:\n%v\n", message, got, want)
	}
}

func randomVec3() mgl32.Vec3 {
	return mgl32.Vec3{
		rand.Float32(), rand.Float32(), rand.Float32(),
	}
}
