package solver

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/generator"
)

type args struct {
	lab  common.Labyrinth
	from common.Location
	to   common.Location
}

type test struct {
	name string
	args args
	want []common.Location
}

func getTestCases() []test {
	tests := []test{
		{
			name: "unconnected maze - can't find way",
			args: args{
				lab:  common.NewLabyrinth(common.NewLocation(1, 0, 0)),
				from: common.NewLocation(1, 0, 0),
				to:   common.NewLocation(0, 0, 0),
			},
			want: nil,
		},

		{
			name: "from out of bounds can't find way",
			args: args{
				lab:  common.NewLabyrinth(common.NewLocation(0, 0, 0)),
				from: common.NewLocation(1, 0, 0),
				to:   common.NewLocation(0, 0, 0),
			},
			want: nil,
		},

		{
			name: "to out of bounds can't find way",
			args: args{
				lab:  common.NewLabyrinth(common.NewLocation(0, 0, 0)),
				from: common.NewLocation(0, 0, 0),
				to:   common.NewLocation(1, 0, 0),
			},
			want: nil,
		},
	}

	{
		lab := common.NewLabyrinth(common.NewLocation(0, 0, 0))
		path := []common.Location{common.NewLocation(0, 0, 0)}

		tests = append(tests, getTestCase("find self", lab, path))
	}

	{
		lab := common.NewLabyrinth(common.NewLocation(1, 0, 0))
		lab.Connect(common.NewLocation(1, 0, 0), common.NewLocation(0, 0, 0))
		path := []common.Location{common.NewLocation(1, 0, 0), common.NewLocation(0, 0, 0)}

		tests = append(tests, getTestCase("one possible stop required - left", lab, path))
	}

	{
		lab := common.NewLabyrinth(common.NewLocation(1, 0, 0))
		lab.Connect(common.NewLocation(1, 0, 0), common.NewLocation(0, 0, 0))
		path := []common.Location{common.NewLocation(0, 0, 0), common.NewLocation(1, 0, 0)}

		tests = append(tests, getTestCase("one possible stop required - right", lab, path))
	}

	{
		lab := common.NewLabyrinth(common.NewLocation(0, 1, 0))
		lab.Connect(common.NewLocation(0, 1, 0), common.NewLocation(0, 0, 0))
		path := []common.Location{common.NewLocation(0, 0, 0), common.NewLocation(0, 1, 0)}

		tests = append(tests, getTestCase("one possible stop required - up", lab, path))
	}

	{
		lab := common.NewLabyrinth(common.NewLocation(0, 1, 0))
		lab.Connect(common.NewLocation(0, 1, 0), common.NewLocation(0, 0, 0))
		path := []common.Location{common.NewLocation(0, 1, 0), common.NewLocation(0, 0, 0)}

		tests = append(tests, getTestCase("one possible stop required - down", lab, path))
	}

	{
		lab := common.NewLabyrinth(common.NewLocation(5, 0, 0))

		lab.Connect(common.NewLocation(0, 0, 0), common.NewLocation(1, 0, 0))
		lab.Connect(common.NewLocation(1, 0, 0), common.NewLocation(2, 0, 0))
		lab.Connect(common.NewLocation(2, 0, 0), common.NewLocation(3, 0, 0))
		lab.Connect(common.NewLocation(3, 0, 0), common.NewLocation(4, 0, 0))
		lab.Connect(common.NewLocation(4, 0, 0), common.NewLocation(5, 0, 0))

		path := []common.Location{
			common.NewLocation(0, 0, 0),
			common.NewLocation(1, 0, 0),
			common.NewLocation(2, 0, 0),
			common.NewLocation(3, 0, 0),
			common.NewLocation(4, 0, 0),
			common.NewLocation(5, 0, 0),
		}

		tests = append(tests, getTestCase("long without branch", lab, path))
	}

	{
		lab := creatStar()

		path := []common.Location{
			common.NewLocation(1, 1, 1),
			common.NewLocation(1, 1, 2),
		}

		tests = append(tests, getTestCase("Star - middle to top", lab, path))
	}

	{
		lab := creatStar()

		path := []common.Location{
			common.NewLocation(1, 1, 1),
			common.NewLocation(1, 1, 0),
		}

		tests = append(tests, getTestCase("Star - middle to bottom", lab, path))
	}

	{
		lab := creatStar()

		path := []common.Location{
			common.NewLocation(1, 1, 1),
			common.NewLocation(1, 0, 1),
		}

		tests = append(tests, getTestCase("Star - middle to front", lab, path))
	}

	{
		lab := creatStar()

		path := []common.Location{
			common.NewLocation(1, 1, 1),
			common.NewLocation(1, 2, 1),
		}

		tests = append(tests, getTestCase("Star - middle to back", lab, path))
	}

	{
		lab := creatStar()

		path := []common.Location{
			common.NewLocation(1, 1, 1),
			common.NewLocation(2, 1, 1),
		}

		tests = append(tests, getTestCase("Star - middle to right", lab, path))
	}

	{
		lab := creatStar()

		path := []common.Location{
			common.NewLocation(1, 1, 1),
			common.NewLocation(0, 1, 1),
		}

		tests = append(tests, getTestCase("Star - middle to left", lab, path))
	}

	return tests
}

func getTestCase(name string, lab common.Labyrinth, path []common.Location) test {
	return test{
		name: name,
		args: args{
			lab:  lab,
			from: path[0],
			to:   path[len(path)-1],
		},
		want: path,
	}
}

func creatStar() common.Labyrinth {
	lab := common.NewLabyrinth(common.NewLocation(2, 2, 2))

	lab.Connect(common.NewLocation(1, 1, 1), common.NewLocation(1, 1, 0))
	lab.Connect(common.NewLocation(1, 1, 1), common.NewLocation(1, 1, 2))
	lab.Connect(common.NewLocation(1, 1, 1), common.NewLocation(1, 0, 1))
	lab.Connect(common.NewLocation(1, 1, 1), common.NewLocation(1, 2, 1))
	lab.Connect(common.NewLocation(1, 1, 1), common.NewLocation(0, 1, 1))
	lab.Connect(common.NewLocation(1, 1, 1), common.NewLocation(2, 1, 1))

	return lab
}

///////////////////    TEST    /////////

func TestRecursiveSolverNoTrust(t *testing.T) {
	for _, tc := range getTestCases() {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := RecursiveSolver(tc.args.lab, tc.args.from, tc.args.to, false); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("RecursiveSolver() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRecursiveSolverWithTrust(t *testing.T) {
	for _, tc := range getTestCases() {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := RecursiveSolver(tc.args.lab, tc.args.from, tc.args.to, true); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("RecursiveSolver() = %v, want %v", got, tc.want)
			}
		})
	}
}

///////////// Benchmarks ///////

func BenchmarkRecursiveSolverWithTrust(b *testing.B) {
	rand.Seed(0)

	lab := generator.NewDepthFirstGenerator().GenerateLabyrinth(common.NewLocation(uint(10), uint(10), uint(b.N)))
	b.ResetTimer()
	RecursiveSolver(lab, common.NewLocation(0, 0, 0), lab.GetMaxLocation(), true)
}

func BenchmarkRecursiveSolverNoTrust(b *testing.B) {
	rand.Seed(0)

	lab := generator.NewDepthFirstGenerator().GenerateLabyrinth(common.NewLocation(uint(10), uint(10), uint(b.N)))
	b.ResetTimer()
	RecursiveSolver(lab, common.NewLocation(0, 0, 0), lab.GetMaxLocation(), true)
}

// Al implementations can (ans should) be tested against all universal test by coping the lines above
// Only the function name and the function in the if statement need to be changed
