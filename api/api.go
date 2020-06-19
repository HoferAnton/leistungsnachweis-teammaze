package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/generator"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/printer"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/solver"
)

const maxRandLabSize uint = 10
const minRandLabSize uint = 3

func MazeAPIRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", mainHandler)
	r.HandleFunc("/{generator}/{solver}", mainHandler)
	r.HandleFunc("/{X:[0-9]+}/{Y:[0-9]+}/{Z:[0-9]+}/{generator}/{solver}", mainHandler)
	r.HandleFunc("/{X:[0-9]+}/{Y:[0-9]+}/{Z:[0-9]+}", mainHandler)

	return r
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var gen generator.LabGenerator

	var solve solver.Function

	switch vars["generator"] {
	case "":
		gen = generator.NewDepthFirstGenerator()
	case "DepthFirstGenerator":
		gen = generator.NewDepthFirstGenerator()
	case "BreadthFirstGenerator":
		gen = generator.NewBreadthFirstGenerator()
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid generator", vars["generator"])

		return
	}

	switch vars["solver"] {
	case "":
		solve = solver.RecursiveSolverSteps
	case "RecursiveSolver":
		solve = solver.RecursiveSolverSteps
	case "ConcurrentSolver":
		solve = solver.ConcurrentSolverSteps
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is not a valid solver", vars["solver"])

		return
	}

	x := stringToUintOrRandom(vars["X"])
	y := stringToUintOrRandom(vars["Y"])
	z := stringToUintOrRandom(vars["Z"])

	fp := common.NewLocation(x, y, z)

	lab, _ := gen.GenerateLabyrinth(fp)
	path, _ := solve(lab, common.NewLocation(0, 0, 0), fp, true)

	w.WriteHeader(http.StatusOK)

	s, _ := printer.Print2D(lab, path)

	fmt.Fprintf(w, "\n%v\n", s)
}

func stringToUintOrRandom(s string) uint {
	u, err := strconv.ParseUint(s, 10, 32)
	if err == nil {
		return uint(u)
	}

	rand.Seed(time.Now().UnixNano())

	return minRandLabSize + uint(rand.Intn(int(maxRandLabSize-minRandLabSize)))
}
