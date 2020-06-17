package display

import (
	"reflect"
	"testing"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/generator"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

func TestMappingColorConverter_ColorMap(t1 *testing.T) {
	type fields struct {
		mapping map[string]mgl32.Vec4
	}

	tests := []struct {
		name   string
		fields fields
		want   map[string]mgl32.Vec4
	}{
		{
			name:   "mapping nil",
			fields: fields{mapping: nil},
			want:   nil,
		},
		{
			name:   "mapping empty",
			fields: fields{mapping: map[string]mgl32.Vec4{}},
			want:   map[string]mgl32.Vec4{},
		},
		{
			name:   "mapping one member",
			fields: fields{mapping: map[string]mgl32.Vec4{"ADD": {0, 1, 0, 1}}},
			want:   map[string]mgl32.Vec4{"ADD": {0, 1, 0, 1}},
		},
		{
			name:   "mapping several members",
			fields: fields{mapping: testMapping()},
			want:   testMapping(),
		},
	}

	for _, tt := range tests {
		tt := tt
		t1.Run(tt.name, func(t1 *testing.T) {
			t := MappingColorConverter{
				mapping: tt.fields.mapping,
			}
			if got := t.ColorMap(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("ColorMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testCubes() []Cube {
	return []Cube{
		testingCubeConstructor(0, 0, 0, 0.5, 0.5, 0.5),
		testingCubeConstructor(1, 0, 0, 0.5, 0.5, 0.5),
		testingCubeConstructor(2, 0, 0, 0.5, 0.5, 0.5),
		testingCubeConstructor(0, 1, 0, 0.5, 0.5, 0.5),
		testingCubeConstructor(1, 1, 0, 0.5, 0.5, 0.5),
		testingCubeConstructor(2, 1, 0, 0.5, 0.5, 0.5),
		testingCubeConstructor(0, 0, 1, 0.5, 0.5, 0.5),
		testingCubeConstructor(1, 0, 1, 0.5, 0.5, 0.5),
		testingCubeConstructor(2, 0, 1, 0.5, 0.5, 0.5),
		testingCubeConstructor(0, 1, 1, 0.5, 0.5, 0.5),
		testingCubeConstructor(1, 1, 1, 0.5, 0.5, 0.5),
		testingCubeConstructor(2, 1, 1, 0.5, 0.5, 0.5),
	}
}

func testLocations() []common.Location {
	return []common.Location{
		common.NewLocation(0, 0, 0),
		common.NewLocation(1, 0, 0),
		common.NewLocation(2, 0, 0),
		common.NewLocation(0, 1, 0),
		common.NewLocation(1, 1, 0),
		common.NewLocation(2, 1, 0),
		common.NewLocation(0, 0, 1),
		common.NewLocation(1, 0, 1),
		common.NewLocation(2, 0, 1),
		common.NewLocation(0, 1, 1),
		common.NewLocation(1, 1, 1),
		common.NewLocation(2, 1, 1),
	}
}

func testLab() *common.Labyrinth {
	lab, _ := generator.NewDepthFirstGenerator().GenerateLabyrinth(common.NewLocation(5, 5, 5))
	return &lab
}

func testingVisualizer() *LabyrinthVisualizer {
	vis := NewLabyrinthVisualizer(testLab(), testingCubeConstructor)
	return &vis
}

func testMapping() map[string]mgl32.Vec4 {
	return map[string]mgl32.Vec4{
		"ADD":    {0, 1, 0, 1},
		"REMOVE": {1, 1, 1, 1},
		"TEST":   {0, 0, 0, 0},
	}
}

func TestMappingColorConverter_StepToColor(t1 *testing.T) {
	type fields struct {
		mapping map[string]mgl32.Vec4
	}

	type args struct {
		step common.Pair
		vis  *LabyrinthVisualizer
	}

	vis := testingVisualizer()

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Cube
		want1  mgl32.Vec4
	}{
		{
			name:   "mapping nil",
			fields: fields{mapping: nil},
			args: args{
				step: common.NewPair(testLocations()[2], "ADD"),
				vis:  vis,
			},
			want:  vis.GetCubeAt(mgl32.Vec3{2, 0, 0}),
			want1: mgl32.Vec4{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t1.Run(tt.name, func(t1 *testing.T) {
			t := MappingColorConverter{
				mapping: tt.fields.mapping,
			}
			got, got1 := t.StepToColor(tt.args.step, tt.args.vis)
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("StepToColor() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t1.Errorf("StepToColor() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMappingColorConverter_PanicTests(t *testing.T) {
	const argumentNilMessage = "nil not allowed in arguments"

	testFunctions := []func(){
		func() {
			want := argumentNilMessage
			defer func() {
				if r := recover(); r != nil {
					if r != want {
						t.Errorf("Unexpected Panic: %v, expected %v", r, want)
					}
				} else {
					t.Errorf("Expected Panic: %v, but none occurred", want)
				}
			}()

			MappingColorConverter{mapping: testMapping()}.
				StepToColor(common.NewPair(testCubes()[0], nil), testingVisualizer())
		},
		func() {
			want := argumentNilMessage
			defer func() {
				if r := recover(); r != nil {
					if r != want {
						t.Errorf("Unexpected Panic: %v, expected %v", r, want)
					}
				} else {
					t.Errorf("Expected Panic: %v, but none occurred", want)
				}
			}()

			MappingColorConverter{mapping: testMapping()}.
				StepToColor(common.NewPair(nil, "VALID"), testingVisualizer())
		},
		func() {
			want := argumentNilMessage
			defer func() {
				if r := recover(); r != nil {
					if r != want {
						t.Errorf("Unexpected Panic: %v, expected %v", r, want)
					}
				} else {
					t.Errorf("Expected Panic: %v, but none occurred", want)
				}
			}()

			MappingColorConverter{mapping: testMapping()}.
				StepToColor(common.NewPair(testLocations()[0], "VALID"), nil)
		},
	}

	for _, f := range testFunctions {
		f()
	}
}
