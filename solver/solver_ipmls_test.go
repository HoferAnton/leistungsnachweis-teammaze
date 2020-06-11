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

func stepsToPath(steps []common.Pair, t *testing.T) []common.Location {
	path := make([]common.Location, 0)

	for _, step := range steps {
		loc := step.GetFirst().(common.Location)
		action := step.GetSecond().(string)

		if action == Add {
			if contains(path, loc) {
				t.Errorf("Added already existig location to path")
			}

			path = append(path, loc)
		} else {
			if !contains(path, loc) {
				t.Errorf("Removed none existing location from path")
			}

			path = removeFirstOccurrence(path, loc)

			if contains(path, loc) {
				t.Errorf("Path must have contained loc more than once, this should not happen")
			}
		}
	}

	if len(path) == 0 {
		path = nil
	}

	return path
}

func pathInSteps(path []common.Location, steps []common.Pair) bool {
	for _, loc := range path {
		found := false

		for _, step := range steps {
			stepLoc := step.GetFirst().(common.Location)
			if loc.Compare(stepLoc) {
				found = true
			}
		}

		if !found {
			return false
		}
	}

	return true
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

func TestConcurrentSolverNoTrust(t *testing.T) {
	for _, tc := range getTestCases() {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := ConcurrentSolver(tc.args.lab, tc.args.from, tc.args.to, false); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ConcurrentSolver() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestConcurrentSolverWithTrust(t *testing.T) {
	for _, tc := range getTestCases() {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := ConcurrentSolver(tc.args.lab, tc.args.from, tc.args.to, true); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ConcurrentSolver() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRecursiveSolverNoTrustWithSteps(t *testing.T) {
	for _, tc := range getTestCases() {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			path, steps := RecursiveSolverSteps(tc.args.lab, tc.args.from, tc.args.to, false)
			if !reflect.DeepEqual(path, tc.want) {
				t.Errorf("RecursiveSolver() = %v, want %v", path, tc.want)
			}
			constructPath := stepsToPath(steps, t)
			if !reflect.DeepEqual(constructPath, tc.want) {
				t.Errorf("RecursiveSolver steps -> constructPath = %v, want %v \nSteps: %v", constructPath, tc.want, steps)
			}
		})
	}
}

func TestRecursiveSolverWithTrustWithSteps(t *testing.T) {
	for _, tc := range getTestCases() {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			path, steps := RecursiveSolverSteps(tc.args.lab, tc.args.from, tc.args.to, true)
			if !reflect.DeepEqual(path, tc.want) {
				t.Errorf("RecursiveSolver() = %v, want %v", path, tc.want)
			}
			constructPath := stepsToPath(steps, t)
			if !reflect.DeepEqual(constructPath, tc.want) {
				t.Errorf("RecursiveSolver steps -> constructPath = %v, want %v \nSteps: %v", constructPath, tc.want, steps)
			}
		})
	}
}

func TestConcurrentSolverNoTrustWithSteps(t *testing.T) {
	for _, tc := range getTestCases() {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			path, steps := ConcurrentSolverSteps(tc.args.lab, tc.args.from, tc.args.to, false)
			if !reflect.DeepEqual(path, tc.want) {
				t.Errorf("ConcurrentSolver() = %v, want %v", path, tc.want)
			}
			if !pathInSteps(path, steps) {
				t.Errorf("ConcurrentSolver path not in steps\npath:\n%v\nsteps:\n%v\n", path, steps)
			}
		})
	}
}

func TestConcurrentSolverWithTrustWithSteps(t *testing.T) {
	for _, tc := range getTestCases() {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			path, steps := ConcurrentSolverSteps(tc.args.lab, tc.args.from, tc.args.to, true)
			if !reflect.DeepEqual(path, tc.want) {
				t.Errorf("ConcurrentSolver() = %v, want %v", path, tc.want)
			}
			if !pathInSteps(path, steps) {
				t.Errorf("ConcurrentSolver path not in steps\npath:\n%v\nsteps:\n%v\n", path, steps)
			}
		})
	}
}

///////////// Benchmarks ///////

func BenchmarkRecursiveSolverWithTrust(b *testing.B) {
	rand.Seed(0)

	lab, _ := generator.NewDepthFirstGenerator().GenerateLabyrinth(common.NewLocation(uint(10), uint(10), uint(10)))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RecursiveSolver(lab, common.NewLocation(0, 0, 0), lab.GetMaxLocation(), true)
	}
}

func BenchmarkRecursiveSolverNoTrust(b *testing.B) {
	rand.Seed(0)

	lab, _ := generator.NewDepthFirstGenerator().GenerateLabyrinth(common.NewLocation(uint(10), uint(10), uint(10)))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RecursiveSolver(lab, common.NewLocation(0, 0, 0), lab.GetMaxLocation(), false)
	}
}

func BenchmarkRecursiveSolverNoTrustWithSteps(b *testing.B) {
	rand.Seed(0)

	lab, _ := generator.NewDepthFirstGenerator().GenerateLabyrinth(common.NewLocation(uint(10), uint(10), uint(10)))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RecursiveSolverSteps(lab, common.NewLocation(0, 0, 0), lab.GetMaxLocation(), false)
	}
}

func BenchmarkRecursiveSolverWithTrustWithSteps(b *testing.B) {
	rand.Seed(0)

	lab, _ := generator.NewDepthFirstGenerator().GenerateLabyrinth(common.NewLocation(uint(10), uint(10), uint(10)))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RecursiveSolverSteps(lab, common.NewLocation(0, 0, 0), lab.GetMaxLocation(), true)
	}
}

func BenchmarkConcurrentSolverWithTrust(b *testing.B) {
	rand.Seed(0)

	lab, _ := generator.NewDepthFirstGenerator().GenerateLabyrinth(common.NewLocation(uint(10), uint(10), uint(10)))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ConcurrentSolver(lab, common.NewLocation(0, 0, 0), lab.GetMaxLocation(), true)
	}
}

func BenchmarkConcurrentSolverNoTrust(b *testing.B) {
	rand.Seed(0)

	lab, _ := generator.NewDepthFirstGenerator().GenerateLabyrinth(common.NewLocation(uint(10), uint(10), uint(10)))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ConcurrentSolver(lab, common.NewLocation(0, 0, 0), lab.GetMaxLocation(), false)
	}
}

func BenchmarkConcurrentSolverNoTrustWithSteps(b *testing.B) {
	rand.Seed(0)

	lab, _ := generator.NewDepthFirstGenerator().GenerateLabyrinth(common.NewLocation(uint(10), uint(10), uint(10)))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ConcurrentSolverSteps(lab, common.NewLocation(0, 0, 0), lab.GetMaxLocation(), false)
	}
}

func BenchmarkConcurrentSolverWithTrustWithSteps(b *testing.B) {
	rand.Seed(0)

	lab, _ := generator.NewDepthFirstGenerator().GenerateLabyrinth(common.NewLocation(uint(10), uint(10), uint(10)))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ConcurrentSolverSteps(lab, common.NewLocation(0, 0, 0), lab.GetMaxLocation(), true)
	}
}
