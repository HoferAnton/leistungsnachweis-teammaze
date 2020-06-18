package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
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

func Run() {
	fmt.Printf("Hi from API")

	r := mux.NewRouter()
	r.HandleFunc("/", Handler)
	r.HandleFunc("/{generator}/{solver}", Handler)
	r.HandleFunc("/{X:[0-9]+}/{Y:[0-9]+}/{Z:[0-9]+}/{generator}/{solver}", Handler)
	r.HandleFunc("/{X:[0-9]+}/{Y:[0-9]+}/{Z:[0-9]+}", Handler)

	http.Handle("/", r)
	_ = http.ListenAndServe(os.Getenv("ADDRESS"), nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
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

	fp, err := maxPoint(vars)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Could not parse uint valuse")
	}

	lab, _ := gen.GenerateLabyrinth(fp)
	path, _ := solve(lab, common.NewLocation(0, 0, 0), fp, true)

	w.WriteHeader(http.StatusOK)

	s, _ := printer.Print2D(lab, path)

	fmt.Fprintf(w, "out: \n%v\n", s)
}

func maxPoint(v map[string]string) (common.Location, error) {
	rand.Seed(time.Now().UnixNano())

	randInt := func() uint { return minRandLabSize + uint(rand.Intn(int(maxRandLabSize-minRandLabSize))) }

	var x, y, z uint

	var err error

	if v["X"] == "" {
		x = randInt()
	} else {
		x, err = stringToUint(v["X"])
		if err != nil {
			return nil, err
		}
	}

	if v["Y"] == "" {
		y = randInt()
	} else {
		y, err = stringToUint(v["Y"])
		if err != nil {
			return nil, err
		}
	}

	if v["Z"] == "" {
		z = randInt()
	} else {
		z, err = stringToUint(v["Z"])
		if err != nil {
			return nil, err
		}
	}

	return common.NewLocation(x, y, z), nil
}

func stringToUint(s string) (uint, error) {
	u, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(u), nil
}
